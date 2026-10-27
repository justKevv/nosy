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

	screen      screen
	drivers     []driver
	connections []connection
	recent      []connection
	picked      int
	connFilter  string
	driver      int
	host        string
	port        string
	database    string
	user        string
	ssl         bool
}

func newApp() *app {
	a := &app{tables: sampleTables(), views: sampleViews(), tabs: sampleTabs(), selected: "users", row: -1,
		screen: screenWorkspace, drivers: sampleDrivers(), connections: sampleConnections(), recent: sampleRecent(), ssl: true}
	a.rows.Selected = &a.row
	a.pickSaved(0)
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
			if a.screen == screenStart {
				a.connectionsSidebar(c)
				verticalRule(c)
				a.connectColumn(c)
				verticalRule(c)
				a.connectionPreview(c)
				return
			}
			a.sidebar(c)
			a.workspace(c)
			a.inspector(c)
		})
		if a.screen == screenStart {
			ui.Divider(c)
			a.startStatusBar(c)
		}
	})
}
