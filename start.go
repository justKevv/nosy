package main

import (
	"strings"

	"github.com/egoist/mygo/ui"
)

// appVersion is the version the status bar of the start screen shows.
const appVersion = "v2.4.1"

// startStatusBar draws the keyboard shortcuts on the left and the count of
// connections, the theme switch and the version on the right.
func (a *app) startStatusBar(c *ui.Context) {
	t := c.Theme()
	ui.Row(c).Height(40).Padding(0, 16).Gap(16).AlignItems(ui.Center).Background(t.Surface).Children(func() {
		shortcut(c, "⌘K", "command palette")
		shortcut(c, "⌘N", "new connection")
		ui.Spacer(c)
		ui.Textf(c, "%d connections", len(a.connections)).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
		ui.Row(c).Width(1).Height(12).Background(t.Border)
		ui.Row(c).Padding(2).Gap(2).Radius(5).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
			ui.Row(c).Padding(3).Radius(4).Children(func() {
				ui.Icon(c, sunIcon).Size(11, 11).TextColor(t.TextMuted)
			})
			ui.Row(c).Padding(3).Radius(4).Background(t.SurfaceHover).Children(func() {
				ui.Icon(c, moonIcon).Size(11, 11).TextColor(t.Text)
			})
		})
		ui.Text(c, appVersion).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
	})
}

// shortcut draws a keyboard key beside what it does.
func shortcut(c *ui.Context, key, label string) {
	t := c.Theme()
	ui.Row(c).Gap(5).AlignItems(ui.Center).Children(func() {
		ui.Row(c).Padding(2, 6).Radius(3).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
			ui.Text(c, key).Font(monoFamily).FontSize(9).TextColor(t.TextMuted).SingleLine()
		})
		ui.Text(c, label).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
	})
}

// pickDriver chooses driver i for the form. Its port is left to the hint.
func (a *app) pickDriver(i int) {
	a.driver = i
	a.port = ""
}

// pickSaved loads saved connection i into the form.
func (a *app) pickSaved(i int) {
	a.picked = i
	a.load(a.connections[i])
}

// pickRecent loads a recently used connection into the form.
func (a *app) pickRecent(conn connection) {
	a.picked = -1
	a.load(conn)
}

// load fills the form with the values of conn.
func (a *app) load(conn connection) {
	a.driver = driverIndex(a.drivers, conn.driver)
	a.host, a.port, a.database, a.user = conn.host, conn.port, conn.database, conn.user
}

// newConnection clears the form for a connection that is not saved yet. Its
// fields show hints: the default host and the port of the driver.
func (a *app) newConnection() {
	a.picked = -1
	a.driver = 0
	a.host, a.port, a.database, a.user = "", "", "", ""
	a.ssl = true
}

// connect opens the workspace of the database the form describes, which
// becomes the target the top bar shows.
func (a *app) connect() {
	name := a.host
	if a.picked >= 0 {
		name = a.connections[a.picked].name
	}
	if name == "" {
		name = "new"
	}
	a.target = &connection{name: name, driver: a.drivers[a.driver].name, host: a.effectiveHost(), port: a.effectivePort(),
		database: a.database, user: a.user, status: "ok"}
	a.screen = screenWorkspace
}

// defaultHost is the host a connection points at when its field is empty.
const defaultHost = "localhost"

// effectiveHost returns the host the form connects to: the one typed, or the
// default host.
func (a *app) effectiveHost() string {
	if a.host != "" {
		return a.host
	}
	return defaultHost
}

// effectivePort returns the port the form connects to: the one typed, or the
// default port of the driver.
func (a *app) effectivePort() string {
	if a.port != "" {
		return a.port
	}
	return a.drivers[a.driver].port
}

// driverOf returns the driver of conn.
func (a *app) driverOf(conn connection) driver {
	return a.drivers[driverIndex(a.drivers, conn.driver)]
}

// uriLines splits the connection string the form describes into its scheme
// and user, its host and port, and its database.
func (a *app) uriLines() []string {
	scheme := a.drivers[a.driver].scheme + "://"
	if a.user != "" {
		scheme += a.user + "@"
	}
	host := a.effectiveHost()
	if port := a.effectivePort(); port != "" {
		host += ":" + port
	}
	lines := []string{scheme, host}
	if a.database != "" {
		lines = append(lines, "/"+a.database)
	}
	return lines
}

// driverIndex returns the index of the driver called name, or 0 if none is.
func driverIndex(drivers []driver, name string) int {
	for i, d := range drivers {
		if d.name == name {
			return i
		}
	}
	return 0
}

// matchesFilter reports whether the name of conn holds filter, ignoring case.
func matchesFilter(conn connection, filter string) bool {
	return strings.Contains(strings.ToLower(conn.name), strings.ToLower(filter))
}

// verticalRule draws the line between two columns of the screen.
func verticalRule(c *ui.Context) {
	ui.Row(c).Width(1).Background(c.Theme().Border)
}

// capsLabel draws the small capital heading of a section of the screen.
func capsLabel(c *ui.Context, text string) {
	ui.Text(c, text).Font(monoFamily).FontSize(10).FontWeight(600).LetterSpacing(0.8).TextColor(c.Theme().TextMuted).SingleLine()
}

// statusDot draws the dot of a connection's status: ok, err, or anything
// else for idle.
func statusDot(c *ui.Context, status string) {
	t := c.Theme()
	color := t.TextMuted
	switch status {
	case "ok":
		color = t.Success
	case "err":
		color = t.Danger
	}
	ui.Row(c).Size(6, 6).Radius(3).Background(color)
}

// inputField draws a text field with an icon, bound to value.
func inputField(c *ui.Context, icon *ui.SVG, value *string, placeholder string) {
	t := c.Theme()
	ui.Row(c).Height(40).Padding(0, 11).Gap(8).Radius(7).Border(1, t.Border).Background(t.Background).AlignItems(ui.Center).Children(func() {
		ui.Icon(c, icon).Size(14, 14).TextColor(t.TextMuted)
		ui.TextInput(c, value).Placeholder(placeholder).Font(monoFamily).Border(0, ui.Transparent).Background(ui.Transparent).Grow(1)
	})
}

// testButton draws the button that tests the form's connection.
func testButton(c *ui.Context) ui.Element {
	t := c.Theme()
	return ui.ButtonBase(c).Padding(7, 14).Gap(6).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
		ui.Icon(c, activityIcon).Size(12, 12).TextColor(t.TextMuted)
		ui.Text(c, "Test").SingleLine()
	})
}

// connectButton draws the accent button that connects the form.
func connectButton(c *ui.Context, connect func()) ui.Element {
	t := c.Theme()
	return ui.ButtonBase(c).Padding(7, 14).Gap(6).Radius(6).Background(t.Accent).AlignItems(ui.Center).Justify(ui.Center).OnClick(connect).Children(func() {
		ui.Text(c, "Connect").FontWeight(600).TextColor(t.AccentText).SingleLine()
		ui.Icon(c, arrowRightIcon).Size(12, 12).TextColor(t.AccentText)
	})
}
