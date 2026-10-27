package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/egoist/mygo/ui"
)

// pageSize is how many rows the grid shows on one page.
const pageSize = 50

// The sizes of the row grid, in DIPs: the heights of the header and of a row,
// the column of row numbers, the bounds of a column's width, the advance of a
// character of its monospace text, the padding of a cell on both sides, and the
// room a status pill takes beside its text.
const (
	gridHeaderHeight   = 34
	gridRowHeight      = 38
	gridRowNumberWidth = 38
	gridMinCellWidth   = 48
	gridMinColumnWidth = 96
	gridMaxColumnWidth = 320
	gridCharWidth      = 7.2
	gridCellPadding    = 20
	gridPillPadding    = 28
)

// sqlLines is the query the editor shows until a script is loaded.
var sqlLines = []string{
	"SELECT id, name, email, created_at",
	"FROM users",
	"WHERE created_at > NOW() - INTERVAL '30 days'",
	"ORDER BY created_at DESC",
	"LIMIT 500;",
}

// filterConditions are the conditions the filter bar shows above the rows.
var filterConditions = []string{"name ILIKE '%a%'", "id > 1"}

// workspace draws the center column: the open tabs, the query toolbar and
// editor, the filter bar, the rows of the selected table and the pager.
func (a *app) workspace(c *ui.Context) {
	ui.Column(c).Grow(1).Children(func() {
		a.tabBar(c)
		a.queryToolbar(c)
		a.editor(c)
		a.filterBar(c)
		a.rowGrid(c)
		a.pager(c)
	})
}

// queryToolbar draws the run actions beside the transaction and limit
// settings and the result status of the last query.
func (a *app) queryToolbar(c *ui.Context) {
	t := c.Theme()
	ui.Row(c).Height(44).Padding(0, 12).Gap(8).AlignItems(ui.Center).Children(func() {
		ui.PrimaryButton(c, "Run Selection ⌘↩")
		ui.ButtonBase(c).Padding(6, 10).Gap(6).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
			ui.Icon(c, boltIcon).Size(14, 14).TextColor(t.TextMuted)
			ui.Text(c, "Explain Analyze").SingleLine()
		})
		metaChip(c, "TX: READ COMMITTED")
		metaChip(c, "LIMIT 500")
		ui.Spacer(c)
		ui.Text(c, "Index Scan · idx_users_created").FontSize(11).TextColor(t.TextMuted).SingleLine()
		metaChip(c, "200 OK · 14.2 ms")
	})
}

// metaChip draws a small bordered label beside the query actions.
func metaChip(c *ui.Context, text string) {
	t := c.Theme()
	ui.Row(c).Padding(4, 8).Radius(6).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
		ui.Text(c, text).FontSize(11).TextColor(t.TextMuted).SingleLine()
	})
}

// editor draws the query text with its line numbers in a gutter.
func (a *app) editor(c *ui.Context) {
	t := c.Theme()
	ui.Column(c).Children(func() {
		ui.Divider(c)
		ui.Row(c).Height(150).Padding(10, 0).Children(func() {
			ui.Column(c).Width(44).AlignItems(ui.End).Padding(0, 10, 0, 0).Children(func() {
				for i := range sqlLines {
					ui.Text(c, fmt.Sprint(i+1)).Font("Menlo").FontSize(12).TextColor(t.TextMuted).SingleLine()
				}
			})
			ui.Column(c).Grow(1).Children(func() {
				for _, line := range sqlLines {
					ui.Text(c, line).Font("Menlo").FontSize(12).SingleLine()
				}
			})
		})
		ui.Divider(c)
	})
}

// filterBar draws the WHERE conditions as removable chips, the add action and
// the count of staged changes.
func (a *app) filterBar(c *ui.Context) {
	t := c.Theme()
	ui.Column(c).Children(func() {
		ui.Row(c).Height(44).Padding(0, 12).Gap(8).AlignItems(ui.Center).Children(func() {
			ui.Row(c).Padding(4, 8).Gap(6).Radius(6).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, filterIcon).Size(12, 12).TextColor(t.TextMuted)
				ui.Text(c, "WHERE").FontSize(11).FontWeight(600).TextColor(t.TextMuted).SingleLine()
			})
			for _, condition := range filterConditions {
				ui.Row(c).Padding(4, 6, 4, 10).Gap(8).Radius(6).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
					ui.Text(c, condition).FontSize(12).SingleLine()
					ui.Icon(c, closeIcon).Size(12, 12).TextColor(t.TextMuted)
				})
			}
			ui.ButtonBase(c).Padding(4, 8).Gap(6).Radius(6).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, plusIcon).Size(12, 12).TextColor(t.TextMuted)
				ui.Text(c, "Add condition").FontSize(12).SingleLine()
			})
			ui.Spacer(c)
			ui.Row(c).Padding(4, 8).Gap(6).Radius(6).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
				ui.Text(c, "0 staged").FontSize(11).TextColor(t.TextMuted).SingleLine()
			})
		})
		ui.Divider(c)
	})
}

// rowGrid draws the rows of the selected table as a grid under a header: a
// column of row numbers, then the columns. Every cell has lines on its right
// and bottom edges, and the header stays at the top while the rows scroll.
// Each column is as wide as its content. When the columns are narrower than
// the grid's area, the last one takes the room left; wider grids scroll
// sideways. Numeric columns align right, and status values show as pills.
func (a *app) rowGrid(c *ui.Context) {
	t := c.Theme()
	cur := a.current()
	widths := []float32{gridRowNumberWidth}
	var total float32
	for i := range cur.columns {
		widths = append(widths, gridColumnWidth(cur, i))
	}
	for _, w := range widths {
		total += w
	}
	if room := a.grid.Bounds(c).W; room > total {
		widths[len(widths)-1] += room - total
		total = room
	}
	ui.ScrollHorizontal(c).Grow(1).AlignItems(ui.Stretch).Bind(&a.grid).Children(func() {
		ui.Column(c).Width(total).AlignItems(ui.Stretch).Children(func() {
			ui.Row(c).Height(gridHeaderHeight).Background(t.Surface).AlignItems(ui.Stretch).Children(func() {
				for j, w := range widths {
					gridCell(c, w, gridAlign(cur, j), func() {
						if j > 0 {
							ui.Text(c, cur.columns[j-1].name).Font("Menlo").FontSize(11).FontWeight(600).TextColor(t.TextMuted).SingleLine()
						}
					})
				}
			})
			ui.List(c, &a.rows, len(cur.rows)).Grow(1).AlignItems(ui.Stretch).Rows(func(r ui.ListRow) {
				gridRow(r, cur, widths)
			})
		})
	})
}

// gridRow draws the row r of cur.
func gridRow(r ui.ListRow, cur table, widths []float32) {
	c := r.Context
	t := c.Theme()
	muted := t.TextMuted
	if r.Selected() {
		muted = t.AccentText
	}
	ui.Row(c).Height(gridRowHeight).AlignItems(ui.Stretch).Children(func() {
		gridCell(c, widths[0], ui.Center, func() {
			ui.Text(c, strconv.Itoa(r.Index+1)).Font("Menlo").FontSize(10).TextColor(muted).SingleLine()
		})
		for j := 1; j < len(widths); j++ {
			gridCell(c, widths[j], gridAlign(cur, j), func() {
				value := cur.rows[r.Index][j-1]
				if cur.columns[j-1].name == "status" {
					statusPill(c, value, muted)
					return
				}
				ui.Text(c, value).Font("Menlo").FontSize(11).SingleLine()
			})
		}
	})
}

// gridCell draws a cell of width w around the content fill builds, with lines
// on its right and bottom edges. Cells keep their width, so text that does
// not fit is cut and the grid scrolls sideways.
func gridCell(c *ui.Context, w float32, align ui.Align, fill func()) {
	t := c.Theme()
	ui.Row(c).Width(w).Padding(0, gridCellPadding/2).AlignItems(ui.Center).Justify(align).Shrink(0).DrawOver(func(p *ui.Painter, r ui.Rect) {
		p.Fill(ui.Rect{X: r.X + r.W - 1, Y: r.Y, W: 1, H: r.H}, t.Border, 0)
		p.Fill(ui.Rect{X: r.X, Y: r.Y + r.H - 1, W: r.W, H: 1}, t.Border, 0)
	}).Children(fill)
}

// gridAlign returns how the cells of column j of cur align: row numbers at
// the center, numbers to the right, and the rest to the left.
func gridAlign(cur table, j int) ui.Align {
	switch {
	case j == 0:
		return ui.Center
	case isNumeric(cur.columns[j-1].typ):
		return ui.End
	}
	return ui.Start
}

// gridColumnWidth returns the width of column i of cur: its longest name or
// value in the first page of rows, within the bounds of a column. Status values
// are measured as the capitals of their pills.
func gridColumnWidth(cur table, i int) float32 {
	chars := utf8.RuneCountInString(cur.columns[i].name)
	status := cur.columns[i].name == "status"
	for _, row := range cur.rows[:min(len(cur.rows), pageSize)] {
		value := row[i]
		if status {
			value = strings.ToUpper(value)
		}
		chars = max(chars, utf8.RuneCountInString(value))
	}
	w := float32(chars)*gridCharWidth + gridCellPadding
	if status {
		w += gridPillPadding
	}
	return min(max(w, gridMinColumnWidth), gridMaxColumnWidth)
}

// statusPill draws a status value as a small bordered pill with a dot, in
// capitals, in the muted color of its row.
func statusPill(c *ui.Context, value string, muted ui.Color) {
	t := c.Theme()
	ui.Row(c).Padding(2, 7).Gap(5).Radius(4).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
		ui.Row(c).Size(5, 5).Radius(3).Background(muted)
		ui.Text(c, strings.ToUpper(value)).Font("Menlo").FontSize(10).FontWeight(600).TextColor(muted).SingleLine()
	})
}

// pager draws the row count of the page beside the page buttons and the
// export action.
func (a *app) pager(c *ui.Context) {
	t := c.Theme()
	cur := a.current()
	pages := max(1, (len(cur.rows)+pageSize-1)/pageSize)
	ui.Column(c).Children(func() {
		ui.Divider(c)
		ui.Row(c).Height(44).Padding(0, 12).Gap(6).AlignItems(ui.Center).Children(func() {
			ui.Textf(c, "Showing 1–%d of %d rows", min(len(cur.rows), pageSize), len(cur.rows)).FontSize(12).TextColor(t.TextMuted).SingleLine()
			ui.Spacer(c)
			ui.ButtonBase(c).Padding(4, 10).Radius(6).Border(1, t.Border).Disabled(true).Children(func() {
				ui.Text(c, "Prev").FontSize(12).SingleLine()
			})
			for p := 1; p <= min(pages, 3); p++ {
				pageButton(c, p, p == 1)
			}
			if pages > 3 {
				ui.Text(c, "…").TextColor(t.TextMuted)
				pageButton(c, pages, false)
			}
			ui.ButtonBase(c).Padding(4, 10).Radius(6).Border(1, t.Border).Children(func() {
				ui.Text(c, "Next").FontSize(12).SingleLine()
			})
			ui.ButtonBase(c).Size(28, 28).Radius(6).Border(1, t.Border).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
				ui.Icon(c, downloadIcon).Size(14, 14).TextColor(t.TextMuted)
			})
		})
	})
}

// pageButton draws one numbered page button of the pager, highlighted when
// it is the current page.
func pageButton(c *ui.Context, n int, current bool) {
	t := c.Theme()
	button := ui.ButtonBase(c).Size(28, 28).Radius(6).Border(1, t.Border).AlignItems(ui.Center).Justify(ui.Center)
	if current {
		button.Background(t.Surface)
	}
	button.Children(func() {
		ui.Text(c, fmt.Sprint(n)).FontSize(12).SingleLine()
	})
}
