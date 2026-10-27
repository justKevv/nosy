package main

import (
	"strings"

	"github.com/egoist/mygo/ui"
)

// quickReconnectCount is how many connections the quick reconnect list shows.
const quickReconnectCount = 4

// connectColumn draws the form for a new connection, centered in the middle of
// the screen: the driver, the target, and the connections to reconnect to.
func (a *app) connectColumn(c *ui.Context) {
	ui.Column(c).Grow(1).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
		ui.Column(c).Width(640).Gap(20).Children(func() {
			connectHead(c)
			a.driverPicker(c)
			a.connectForm(c)
			a.quickReconnect(c)
		})
	})
}

// connectHead draws the title of the form and what it does.
func connectHead(c *ui.Context) {
	t := c.Theme()
	ui.Column(c).Gap(4).Children(func() {
		ui.Text(c, "New connection").FontSize(18).FontWeight(600).SingleLine()
		ui.Text(c, "Pick a driver, point it at a target, connect. Drivers load on first use.").FontSize(12).TextColor(t.TextMuted)
	})
}

// driverPicker draws the chips of the drivers, the chosen one highlighted.
func (a *app) driverPicker(c *ui.Context) {
	t := c.Theme()
	ui.Column(c).Gap(9).Children(func() {
		capsLabel(c, "DRIVER")
		ui.Row(c).Gap(6).AlignItems(ui.Center).Wrap().Children(func() {
			for i, d := range a.drivers {
				chip := ui.ButtonBase(c).Padding(7, 9).Gap(6).Radius(6).Border(1, t.Border).AlignItems(ui.Center).OnClick(func() { a.pickDriver(i) })
				if i == a.driver {
					chip = chip.Border(1, t.Accent).Background(t.Surface)
				}
				chip.Children(func() {
					ui.Icon(c, d.icon).Size(13, 13).TextColor(t.TextMuted)
					ui.Text(c, d.name).FontSize(12).SingleLine()
				})
			}
			ui.ButtonBase(c).Padding(7, 9).Gap(6).Radius(6).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, plusIcon).Size(12, 12).TextColor(t.TextMuted)
				ui.Text(c, "More").FontSize(12).TextColor(t.TextMuted).SingleLine()
			})
		})
	})
}

// connectForm draws the fields of the connection: the host and port, the
// database and user, the connection string they make, and the actions.
func (a *app) connectForm(c *ui.Context) {
	t := c.Theme()
	ui.Column(c).Padding(16).Gap(12).Radius(10).Border(1, t.Border).Background(t.Surface).Children(func() {
		ui.Row(c).Gap(12).Children(func() {
			ui.Column(c).Grow(1).Gap(6).Children(func() {
				capsLabel(c, "HOST")
				inputField(c, globeIcon, &a.host, "hostname or IP")
			})
			ui.Column(c).Width(92).Gap(6).Children(func() {
				capsLabel(c, "PORT")
				inputField(c, hashIcon, &a.port, "")
			})
		})
		ui.Row(c).Gap(12).Children(func() {
			ui.Column(c).Grow(1).Gap(6).Children(func() {
				capsLabel(c, "DATABASE")
				inputField(c, databaseIcon, &a.database, "")
			})
			ui.Column(c).Grow(1).Gap(6).Children(func() {
				capsLabel(c, "USER")
				inputField(c, userIcon, &a.user, "")
			})
		})
		ui.Column(c).Gap(6).Children(func() {
			capsLabel(c, "CONNECTION STRING")
			ui.Row(c).Height(40).Padding(0, 11).Gap(8).Radius(7).Border(1, t.Border).Background(t.Background).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, linkIcon).Size(12, 12).TextColor(t.TextMuted)
				ui.Text(c, strings.Join(a.uriLines(), "")).Font(monoFamily).FontSize(11).TextColor(t.TextMuted).SingleLine()
			})
		})
		ui.Row(c).Gap(10).AlignItems(ui.Center).Children(func() {
			ui.Checkbox(c, &a.ssl, "Require SSL")
			ui.Text(c, "timeout 10s").Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
			ui.Spacer(c)
			testButton(c)
			connectButton(c, a.connect)
		})
	})
}

// quickReconnect draws the connections to open again, one row each.
func (a *app) quickReconnect(c *ui.Context) {
	t := c.Theme()
	ui.Column(c).AlignItems(ui.Stretch).Radius(10).Border(1, t.Border).Background(t.Surface).Children(func() {
		ui.Row(c).Padding(11, 14).Gap(8).AlignItems(ui.Center).Children(func() {
			ui.Icon(c, historyIcon).Size(13, 13).TextColor(t.TextMuted)
			capsLabel(c, "QUICK RECONNECT")
			ui.Spacer(c)
			ui.Text(c, "last used · all").Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
			ui.Text(c, "View all").FontSize(11).TextColor(t.Accent).SingleLine()
		})
		for i, conn := range a.connections[:min(quickReconnectCount, len(a.connections))] {
			ui.Divider(c)
			a.quickRow(c, i, conn)
		}
	})
}

// quickRow draws connection i of the quick reconnect list: its status, name,
// driver, database, latency and an arrow to open it.
func (a *app) quickRow(c *ui.Context, i int, conn connection) {
	t := c.Theme()
	ui.ButtonBase(c).Padding(9, 14).Gap(10).AlignItems(ui.Center).OnClick(func() { a.pickSaved(i) }).Children(func() {
		statusDot(c, conn.status)
		ui.Text(c, conn.name).Font(monoFamily).FontSize(11).SingleLine()
		ui.Text(c, conn.driver).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
		ui.Spacer(c)
		ui.Text(c, conn.database).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
		ui.Text(c, conn.latency).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
		ui.Icon(c, arrowUpRightIcon).Size(10, 10).TextColor(t.TextMuted)
	})
}
