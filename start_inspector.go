package main

import (
	"strings"

	"github.com/egoist/mygo/ui"
)

// connectionPreview draws the right column of the start screen: the target
// the form describes, its fields and the connection string it resolves to.
func (a *app) connectionPreview(c *ui.Context) {
	t := c.Theme()
	d := a.drivers[a.driver]
	ui.Column(c).Width(328).AlignItems(ui.Stretch).Children(func() {
		ui.Row(c).Height(40).Padding(0, 14).Gap(8).AlignItems(ui.Center).Background(t.Surface).Children(func() {
			ui.Icon(c, plugIcon).Size(13, 13).TextColor(t.Accent)
			capsLabel(c, "CONNECTION PREVIEW")
			ui.Spacer(c)
			ui.Row(c).Padding(1, 6).Radius(4).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
				ui.Text(c, "read-only").Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
			})
		})
		ui.Divider(c)
		ui.Column(c).Padding(12, 14).Gap(8).Children(func() {
			capsLabel(c, "TARGET")
			ui.Row(c).Gap(10).AlignItems(ui.Center).Children(func() {
				ui.Row(c).Size(30, 30).Radius(7).Border(1, t.Border).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
					ui.Icon(c, d.icon).Size(14, 14).TextColor(t.Accent)
				})
				ui.Column(c).Grow(1).Gap(2).Children(func() {
					ui.Text(c, d.version).FontSize(13).FontWeight(600).SingleLine()
					ui.Text(c, d.wire).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
				})
				ui.Row(c).Padding(2, 7).Gap(4).Radius(4).Border(1, t.Border).AlignItems(ui.Center).Children(func() {
					statusDot(c, "")
					ui.Text(c, "idle").Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
				})
			})
		})
		ui.Divider(c)
		ui.Column(c).Padding(12, 14).Gap(7).Children(func() {
			capsLabel(c, "FIELDS")
			fieldRow(c, "host", orDash(a.host))
			fieldRow(c, "port", orDash(a.port))
			fieldRow(c, "database", orDash(a.database))
			fieldRow(c, "user", orDash(a.user))
			ssl := "off"
			if a.ssl {
				ssl = "required"
			}
			fieldRow(c, "ssl", ssl)
		})
		ui.Divider(c)
		ui.Column(c).Padding(12, 14).Gap(6).Children(func() {
			capsLabel(c, "RESOLVED URI")
			ui.Column(c).Padding(8, 9).Gap(3).Radius(6).Border(1, t.Border).Background(t.Background).Children(func() {
				for _, line := range a.uriLines() {
					ui.Text(c, line).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
				}
			})
		})
		ui.Spacer(c)
		ui.Divider(c)
		ui.Row(c).Padding(8, 14).Gap(8).Children(func() {
			testButton(c).Grow(1)
			connectButton(c, a.connect).Grow(1)
		})
	})
}

// fieldRow draws one field of the preview: its name beside its value.
func fieldRow(c *ui.Context, name, value string) {
	t := c.Theme()
	ui.Row(c).AlignItems(ui.Center).Justify(ui.SpaceBetween).Gap(8).Children(func() {
		ui.Text(c, name).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
		ui.Text(c, value).Font(monoFamily).FontSize(10).SingleLine()
	})
}

// orDash returns value, or a dash for a field the form leaves empty.
func orDash(value string) string {
	if strings.TrimSpace(value) == "" {
		return "—"
	}
	return value
}
