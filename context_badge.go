package main

import "github.com/egoist/mygo/ui"

// contextBadge draws the database the window is on, in the top bar: the
// connection, its database and the engine version. Without a target it asks
// for one.
func (a *app) contextBadge(c *ui.Context) {
	t := c.Theme()
	ui.Row(c).Padding(5, 10).Gap(8).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
		if a.target == nil {
			statusDot(c, "")
			ui.Text(c, "not connected").Font(monoFamily).FontSize(12).TextColor(t.TextMuted).SingleLine()
			ui.Text(c, "/").Font(monoFamily).FontSize(12).TextColor(t.TextMuted)
			ui.Text(c, "choose a database").Font(monoFamily).FontSize(12).FontWeight(600).TextColor(t.Accent).SingleLine()
		} else {
			statusDot(c, a.target.status)
			ui.Text(c, a.target.name).Font(monoFamily).FontSize(12).TextColor(t.Text).SingleLine()
			ui.Text(c, "/").Font(monoFamily).FontSize(12).TextColor(t.TextMuted)
			ui.Text(c, a.target.database).Font(monoFamily).FontSize(12).FontWeight(600).TextColor(t.Accent).SingleLine()
			ui.Row(c).Padding(1, 5).Radius(4).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
				ui.Text(c, a.driverOf(*a.target).tag).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
			})
		}
		ui.Icon(c, chevronIcon).Size(12, 12).TextColor(t.TextMuted)
	})
}
