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
	grid     ui.Handle
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
	ui.Column(c).Fill().Children(func() {
		a.topBar(c)
		ui.Row(c).Grow(1).AlignItems(ui.Stretch).Children(func() {
			a.sidebar(c)
			a.workspace(c)
			a.inspector(c)
		})

	})
}
