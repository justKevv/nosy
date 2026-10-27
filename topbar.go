package main

import "github.com/egoist/mygo/ui"

// titleBarHeight is the height of the top bar. It holds the native window
// controls, which are centered in it.
const titleBarHeight = 44

// topBar draws the window's top row in the title bar, beside the window
// controls, which it leaves room for. Both screens share its layout; the items
// are the ones each screen needs.
func (a *app) topBar(c *ui.Context) {
	t := c.Theme()
	bar := c.TitleBar()
	ui.Row(c).Height(titleBarHeight).Border(1, t.Border).Padding(0, bar.Right+15, 0, bar.Left+20).Gap(10).AlignItems(ui.Center).DragWindow().Children(func() {
		ui.Row(c).Grow(1).Gap(10).AlignItems(ui.Center).Children(func() {
			ui.Row(c).Height(24).AlignItems(ui.Center).Children(func() {
				ui.Divider(c)
			})
			ui.Row(c).Size(24, 24).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
				ui.Icon(c, databaseIcon).Size(14, 14).TextColor(t.Text)
			})
			if a.screen == screenStart {
				ui.Text(c, "nosy").Bold().FontSize(14).TextColor(t.TextMuted)
			} else {
				ui.Text(c, "nosy").Bold().FontSize(14).TextColor(t.TextMuted)
			}
			a.contextBadge(c)
		})
		ui.Row(c).Grow(1).Gap(10).AlignItems(ui.Center).Children(func() {
			ui.Spacer(c)
			if a.screen == screenStart {
				ui.Row(c).Padding(6, 10).Gap(6).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
					ui.Icon(c, layersIcon).Size(12, 12).TextColor(t.TextMuted)
					ui.Textf(c, "%d saved", len(a.connections)).FontSize(11).TextColor(t.TextMuted).SingleLine()
				})
				ui.Row(c).Padding(6, 10).Gap(6).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
					ui.Icon(c, uploadIcon).Size(12, 12).TextColor(t.TextMuted)
					ui.Text(c, "Import").SingleLine()
				})
				ui.PrimaryButton(c, "+ New Connection").OnClick(a.newConnection)
				return
			}
			ui.TextInput(c, &a.search).Placeholder("Jump to table, index, or SQL…").Width(300)
			ui.Row(c).Padding(4, 9).Gap(6).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, activityIcon).Size(12, 12).TextColor(t.TextMuted)
				ui.Text(c, "connected · 1.8ms").TextColor(t.TextMuted).SingleLine()
			})
			ui.PrimaryButton(c, "+ New Query")
		})
	})
}
