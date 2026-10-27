package main

import (
	"strconv"

	"github.com/egoist/mygo/ui"
)

// connectionsSidebar draws the left column of the start screen: the scope and
// the filter, the saved connections, the recent ones and the driver cache.
func (a *app) connectionsSidebar(c *ui.Context) {
	t := c.Theme()
	ui.Column(c).Width(264).AlignItems(ui.Stretch).Children(func() {
		ui.Column(c).Padding(12).Gap(8).Children(func() {
			ui.Row(c).Padding(6, 10).Gap(8).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, folderIcon).Size(13, 13).TextColor(t.TextMuted)
				ui.Text(c, "scope:").Font("Menlo").FontSize(12).TextColor(t.TextMuted)
				ui.Text(c, "all").Font("Menlo").FontSize(12).FontWeight(600).SingleLine()
				ui.Spacer(c)
				ui.Textf(c, "%d saved", len(a.connections)).Font("Menlo").FontSize(10).TextColor(t.TextMuted).SingleLine()
				ui.Icon(c, chevronsIcon).Size(12, 12).TextColor(t.TextMuted)
			})
			inputField(c, filterIcon, &a.connFilter, "Filter connections...")
		})
		ui.Divider(c)
		ui.Column(c).Padding(12, 8, 8, 8).Gap(2).Children(func() {
			ui.Row(c).Padding(4, 6).Gap(6).AlignItems(ui.Center).Children(func() {
				capsLabel(c, "SAVED")
				ui.Badge(c, strconv.Itoa(len(a.connections)))
				ui.Spacer(c)
				ui.ButtonBase(c).Padding(2).Radius(4).OnClick(a.newConnection).Children(func() {
					ui.Icon(c, plusIcon).Size(12, 12).TextColor(t.TextMuted)
				})
			})
			for i, conn := range a.connections {
				if matchesFilter(conn, a.connFilter) {
					a.savedRow(c, i, conn)
				}
			}
		})
		ui.Divider(c)
		ui.Column(c).Padding(12, 8, 4, 8).Gap(2).Children(func() {
			ui.Row(c).Padding(4, 6).Gap(6).AlignItems(ui.Center).Children(func() {
				capsLabel(c, "RECENT")
				ui.Badge(c, "unsaved")
				ui.Spacer(c)
				ui.ButtonBase(c).Padding(2, 4).Radius(4).OnClick(func() { a.recent = nil }).Children(func() {
					ui.Text(c, "Clear").Font("Menlo").FontSize(10).TextColor(t.TextMuted).SingleLine()
				})
			})
			for _, conn := range a.recent {
				if matchesFilter(conn, a.connFilter) {
					a.recentRow(c, conn)
				}
			}
		})
		ui.Spacer(c)
	})
}

// savedRow draws saved connection i of the sidebar, chosen rows in a surface
// with a border.
func (a *app) savedRow(c *ui.Context, i int, conn connection) {
	t := c.Theme()
	row := ui.ButtonBase(c).Padding(6, 8).Gap(10).Radius(6).AlignItems(ui.Center).OnClick(func() { a.pickSaved(i) })
	if i == a.picked {
		row = row.Background(t.Surface).Border(1, t.Border)
	}
	row.Children(func() {
		ui.Row(c).Size(28, 28).Radius(6).Border(1, t.Border).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
			ui.Icon(c, a.driverOf(conn).icon).Size(14, 14).TextColor(t.TextMuted)
		})
		ui.Column(c).Grow(1).Gap(2).Children(func() {
			ui.Text(c, conn.name).Font("Menlo").FontSize(12).SingleLine()
			ui.Text(c, subtitle(conn)).Font("Menlo").FontSize(10).TextColor(t.TextMuted).SingleLine()
		})
		ui.Column(c).AlignItems(ui.End).Gap(2).Children(func() {
			metric, color := conn.latency, t.TextMuted
			if conn.status == "err" {
				metric, color = "err", t.Danger
			}
			ui.Text(c, metric).Font("Menlo").FontSize(10).TextColor(color).SingleLine()
			ui.Text(c, conn.ago).Font("Menlo").FontSize(9).TextColor(t.TextMuted).SingleLine()
		})
	})
}

// recentRow draws a recently used connection of the sidebar.
func (a *app) recentRow(c *ui.Context, conn connection) {
	t := c.Theme()
	ui.ButtonBase(c).Padding(7, 8).Gap(8).Radius(5).AlignItems(ui.Center).OnClick(func() { a.pickRecent(conn) }).Children(func() {
		statusDot(c, conn.status)
		ui.Text(c, conn.name).Font("Menlo").FontSize(12).Grow(1).SingleLine()
		ui.Text(c, conn.ago).Font("Menlo").FontSize(9).TextColor(t.TextMuted).SingleLine()
	})
}

// subtitle is the line under a connection's name: its database and driver.
func subtitle(conn connection) string {
	if conn.database == "" {
		return conn.driver
	}
	return conn.database + " · " + conn.driver
}
