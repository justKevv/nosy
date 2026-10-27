package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// pageSize is how many rows the grid shows on one page.
const pageSize = 50

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

// rowGrid draws the rows of the selected table, one row per line.
func (a *app) rowGrid(c *ui.Context) {
	cur := a.current()
	columns := make([]ui.TableColumn, len(cur.columns))
	for i, col := range cur.columns {
		columns[i] = ui.TableColumn{Title: col.name}
	}
	ui.Table(c, &a.rows, columns, len(cur.rows), func(row, col int) {
		ui.Text(c, cur.rows[row][col]).SingleLine()
	}).Grow(1)
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
