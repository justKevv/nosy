package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// statusBar draws the footer under the window: the tablespace and encoding of
// the connection on the left, the open table and its row count on the right.
func (a *app) statusBar(c *ui.Context) {
	t := c.Theme()
	cur := a.current()
	ui.Column(c).Children(func() {
		ui.Divider(c)
		ui.Row(c).Height(28).Padding(0, 12).Gap(16).AlignItems(ui.Center).Children(func() {
			ui.Text(c, "tablespace: pg_default").FontSize(11).TextColor(t.TextMuted).SingleLine()
			ui.Text(c, "UTF8").FontSize(11).TextColor(t.TextMuted).SingleLine()
			ui.Text(c, "UTC").FontSize(11).TextColor(t.TextMuted).SingleLine()
			ui.Spacer(c)
			ui.Text(c, fmt.Sprintf("public.%s", cur.name)).FontSize(11).TextColor(t.TextMuted).SingleLine()
			ui.Text(c, fmt.Sprintf("%d rows", len(cur.rows))).FontSize(11).TextColor(t.TextMuted).SingleLine()
			ui.Text(c, "ready").FontSize(11).TextColor(t.TextMuted).SingleLine()
		})
	})
}
