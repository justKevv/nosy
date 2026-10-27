package main

import (
	"fmt"
	"strings"

	"github.com/egoist/mygo/ui"
)

// sidebar draws the left column: the schema and the filter above the tables
// and views of the schema, which scroll beneath them.
func (a *app) sidebar(c *ui.Context) {
	t := c.Theme()
	ui.Column(c).Width(240).Border(1, t.Border).Children(func() {
		ui.Column(c).Padding(12, 12, 12, 12).Gap(8).Children(func() {
			ui.Row(c).Padding(7, 10).Gap(8).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, folderIcon).Size(14, 14).TextColor(t.TextMuted)
				ui.Text(c, "schema:").TextColor(t.TextMuted).SingleLine()
				ui.Text(c, "public").Bold().SingleLine()
				ui.Spacer(c)
				ui.Text(c, "18 obj").FontSize(11).TextColor(t.TextMuted).SingleLine()
				ui.Icon(c, chevronsIcon).Size(12, 12).TextColor(t.TextMuted)
			})
			ui.Row(c).Padding(7, 10).Gap(8).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, filterIcon).Size(14, 14).TextColor(t.TextMuted)
				ui.TextInputBase(c, &a.filter).Placeholder("Filter objects…").Grow(1)
				ui.Text(c, "/").TextColor(t.TextMuted)
			})
		})
		ui.Divider(c)
		ui.Sidebar(c, &a.selected, func() {
			sectionHeader(c, fmt.Sprintf("Tables (%d)", len(a.tables)), func() {
				ui.Icon(c, plusIcon).Size(14, 14).TextColor(t.TextMuted)
			})
			for _, tb := range a.tables {
				ui.SidebarItem(c, tb.name, tableIcon, tb.name).Children(func() {
					ui.Text(c, fmt.Sprint(len(tb.rows))).FontSize(12).TextColor(t.TextMuted).SingleLine()
				})
				if tb.name == a.selected {
					for _, col := range tb.columns {
						columnRow(c, col)
					}
				}
			}
			ui.Divider(c).Margin(6, 0, 6, 0)
			sectionHeader(c, fmt.Sprintf("Views & Matviews (%d)", len(a.views)), nil)
			for _, v := range a.views {
				icon := eyeIcon
				if v.kind == "plpgsql" {
					icon = functionIcon
				}
				ui.SidebarItem(c, v.name, icon, v.name).Children(func() {
					ui.Text(c, v.kind).FontSize(12).TextColor(t.TextMuted).SingleLine()
				})
			}
		}).Grow(1).OnChange(func() { a.row = -1 })
	})
}

// sectionHeader draws the muted title of a sidebar section, with trailing
// beside it when trailing is not nil.
func sectionHeader(c *ui.Context, title string, trailing func()) {
	t := c.Theme()
	ui.Row(c).Padding(6, 8).AlignItems(ui.Center).Children(func() {
		ui.Text(c, strings.ToUpper(title)).FontSize(11).FontWeight(600).TextColor(t.TextMuted).Grow(1).SingleLine()
		if trailing != nil {
			trailing()
		}
	})
}

// columnRow draws a column of the selected table under its table in the
// sidebar: its key, foreign key or number icon, its name and its type.
func columnRow(c *ui.Context, col column) {
	t := c.Theme()
	ui.Row(c).Padding(4, 8, 4, 36).Gap(8).AlignItems(ui.Center).Children(func() {
		if icon := columnIcon(col); icon != nil {
			ui.Icon(c, icon).Size(13, 13).TextColor(t.TextMuted)
		} else {
			ui.Row(c).Size(13, 13)
		}
		ui.Text(c, col.name).Grow(1).SingleLine()
		ui.Text(c, col.typ).FontSize(11).TextColor(t.TextMuted).SingleLine()
	})
}

// columnIcon picks the icon that marks a column: a key for the primary key,
// a link for a foreign key and a hash for a number. It is nil for the rest.
func columnIcon(col column) *ui.SVG {
	switch {
	case col.name == "id":
		return keyIcon
	case strings.HasSuffix(col.name, "_id"):
		return linkIcon
	case strings.HasPrefix(col.typ, "numeric") || strings.HasPrefix(col.typ, "bigint"):
		return hashIcon
	}
	return nil
}
