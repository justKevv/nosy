package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// inspectorWidth is the width of the row inspector on the right.
const inspectorWidth = 300

// inspector draws the right column: the values of the selected row, one
// column per block with its name and SQL type above the value.
func (a *app) inspector(c *ui.Context) {
	t := c.Theme()
	cur := a.current()
	selected := a.row >= 0 && a.row < len(cur.rows)
	ui.Column(c).Width(inspectorWidth).Border(1, t.Border).Children(func() {
		ui.Row(c).Height(44).Padding(0, 12).Gap(8).AlignItems(ui.Center).Children(func() {
			ui.Text(c, "Row Inspector").Bold().SingleLine()
			ui.Spacer(c)
			if selected {
				ui.Text(c, fmt.Sprintf("row %d", a.row+1)).FontSize(11).TextColor(t.TextMuted).SingleLine()
			}
		})
		ui.Divider(c)
		if !selected {
			ui.Column(c).Grow(1).Padding(16).Children(func() {
				ui.Text(c, "Select a row to inspect its values.").TextColor(t.TextMuted)
			})
			return
		}
		ui.Column(c).Grow(1).Padding(12).Gap(12).Children(func() {
			for i, col := range cur.columns {
				ui.Column(c).Gap(4).Children(func() {
					ui.Row(c).Gap(6).AlignItems(ui.Center).Children(func() {
						ui.Text(c, col.name).FontSize(11).FontWeight(600).TextColor(t.TextMuted).SingleLine()
						ui.Text(c, col.typ).FontSize(11).TextColor(t.TextMuted).SingleLine()
					})
					ui.Text(c, cur.rows[a.row][i]).Font("Menlo").FontSize(12).SingleLine()
				})
			}
		})
	})
}
