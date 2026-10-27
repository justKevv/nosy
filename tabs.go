package main

import "github.com/egoist/mygo/ui"

// workspaceTab is a tab open above the table grid: a table, a SQL script or
// an ERD diagram. A badge names its kind, and dirty marks unsaved changes.
type workspaceTab struct {
	name  string
	icon  *ui.SVG
	badge string
	dirty bool
}

func sampleTabs() []workspaceTab {
	return []workspaceTab{
		{name: "public.query_executions", icon: tableIcon, badge: "DATA"},
		{name: "slow_locks_audit.sql", icon: terminalIcon, dirty: true},
		{name: "telemetry_core.erd", icon: diagramIcon},
	}
}

// tabBar draws the open tabs beside the Relationships button, over a divider.
// The active tab shows a close mark.
func (a *app) tabBar(c *ui.Context) {
	t := c.Theme()
	ui.Column(c).Children(func() {
		ui.Row(c).Height(40).AlignItems(ui.Center).Gap(8).Padding(0, 8, 0, 0).Children(func() {
			tabs := ui.TabsBase(c, &a.tab, len(a.tabs))
			tabs.List.Grow(1).Height(40).Gap(4).Children(func() {
				for i, tb := range a.tabs {
					tab := tabs.Tab(i).Height(40).Padding(0, 12).Gap(8).AlignItems(ui.Center)
					if i == a.tab {
						tab.Background(t.Surface)
					}
					tab.Children(func() {
						ui.Icon(c, tb.icon).Size(14, 14).TextColor(t.TextMuted)
						ui.Text(c, tb.name).SingleLine()
						if tb.badge != "" {
							ui.Badge(c, tb.badge)
						}
						if tb.dirty {
							ui.Row(c).Size(6, 6).Radius(3).Background(t.TextMuted)
						}
						if i == a.tab {
							ui.Icon(c, closeIcon).Size(12, 12).TextColor(t.TextMuted)
						}
					})
				}
			})
			ui.ButtonBase(c).Padding(6, 10).Gap(6).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, diagramIcon).Size(14, 14).TextColor(t.TextMuted)
				ui.Text(c, "Relationships").SingleLine()
			})
		})
		ui.Divider(c)
	})
}
