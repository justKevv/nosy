package main

import "github.com/egoist/mygo/ui"

// app is the state the window shows. Its view builds the interface from
// it, on the main thread, whenever the window needs a frame: after input,
// after Window.Update, and while something animates.
type app struct {
	tables   []table
	views    []view
	tabs     []workspaceTab
	tab      int
	selected string
	search   string
	filter   string
	rows     ui.ListState
	row      int
}

func newApp() *app {
	a := &app{tables: sampleTables(), views: sampleViews(), tabs: sampleTabs(), selected: "users", row: -1}
	a.rows.Selected = &a.row
	return a
}

func (a *app) current() table {
	for _, t := range a.tables {
		if t.name == a.selected {
			return t
		}
	}
	return table{}
}

func (a *app) view(c *ui.Context) {
	t := c.Theme()
	cur := a.current()
	ui.Column(c).Fill().Children(func() {
		a.topBar(c)
		ui.Row(c).Grow(1).AlignItems(ui.Stretch).Children(func() {
			a.sidebar(c)
			ui.Column(c).Grow(1).Children(func() {
				a.tabBar(c)
				columns := make([]ui.TableColumn, len(cur.columns))
				for i, col := range cur.columns {
					columns[i] = ui.TableColumn{Title: col.name}
				}
				ui.Table(c, &a.rows, columns, len(cur.rows), func(row, col int) {
					ui.Text(c, cur.rows[row][col]).SingleLine()
				}).Grow(1)
				ui.Textf(c, "%d rows · %d columns", len(cur.rows), len(cur.columns)).
					TextColor(t.TextMuted).Padding(6, 12)
			})
		})
	})
}
