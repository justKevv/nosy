package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/egoist/mygo/ui"
)

// inspectorWidth is the width of the row inspector on the right.
const inspectorWidth = 328

// inspector draws the right column in its read-only layout: the header with
// the selected row's number, the columns and types with their values, the
// record as JSON, and the action that copies it.
func (a *app) inspector(c *ui.Context) {
	t := c.Theme()
	cur := a.current()
	selected := a.row >= 0 && a.row < len(cur.rows)
	var copyJSON ui.Element
	ui.Column(c).Width(inspectorWidth).Border(1, t.Border).Children(func() {
		ui.Row(c).Height(44).Padding(0, 14).Gap(8).AlignItems(ui.Center).Children(func() {
			ui.Icon(c, sidebarIcon).Size(14, 14).TextColor(t.TextMuted)
			ui.Text(c, "ROW INSPECTOR").Font(monoFamily).FontSize(11).FontWeight(600).LetterSpacing(0.6).TextColor(t.TextMuted).SingleLine()
			ui.Spacer(c)
			if selected {
				rowBadge(c, fmt.Sprintf("Row #%d", a.row+1))
			}
		})
		ui.Divider(c)
		if !selected {
			ui.Column(c).Grow(1).Padding(16).Children(func() {
				ui.Text(c, "Select a row to inspect its values.").TextColor(t.TextMuted)
			})
		} else {
			ui.Column(c).Grow(1).Children(func() {
				ui.Column(c).Padding(12, 14).Gap(6).Children(func() {
					ui.Row(c).AlignItems(ui.Center).Children(func() {
						ui.Text(c, "COLUMNS & TYPES").Font(monoFamily).FontSize(10).FontWeight(600).LetterSpacing(0.6).TextColor(t.TextMuted).SingleLine()
						ui.Spacer(c)
						ui.Text(c, fmt.Sprint(len(cur.columns))).Font(monoFamily).FontSize(9).TextColor(t.TextMuted).SingleLine()
					})
					for i, col := range cur.columns {
						if i > 0 {
							ui.Divider(c)
						}
						ui.Row(c).Padding(5, 0).Gap(8).AlignItems(ui.Center).Justify(ui.SpaceBetween).Children(func() {
							ui.Column(c).Gap(1).Children(func() {
								ui.Text(c, col.name).Font(monoFamily).FontSize(10).TextColor(t.Text).SingleLine()
								ui.Text(c, col.typ).Font(monoFamily).FontSize(9).TextColor(t.TextMuted).SingleLine()
							})
							ui.Text(c, cur.rows[a.row][i]).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
						})
					}
				})
				ui.Divider(c)
				ui.Column(c).Grow(1).Padding(12, 14).Gap(6).Children(func() {
					ui.Row(c).AlignItems(ui.Center).Gap(6).Children(func() {
						ui.Icon(c, bracesIcon).Size(13, 13).TextColor(t.TextMuted)
						ui.Text(c, "RECORD JSON").Font(monoFamily).FontSize(10).FontWeight(600).LetterSpacing(0.6).TextColor(t.TextMuted).SingleLine()
						ui.Spacer(c)
						ui.Text(c, "read-only").Font(monoFamily).FontSize(9).TextColor(t.TextMuted).SingleLine()
					})
					ui.Column(c).Grow(1).Padding(10).Gap(3).Radius(6).Border(1, t.Border).Background(t.Surface).Children(func() {
						for _, line := range recordLines(cur, a.row) {
							ui.Text(c, line).Font(monoFamily).FontSize(11).SingleLine()
						}
					})
				})
			})
		}
		ui.Divider(c)
		ui.Row(c).Padding(8, 14).Children(func() {
			copyJSON = ui.PrimaryButton(c, "Copy JSON").Grow(1).Disabled(!selected)
		})
	})
	if selected && copyJSON.Clicked() {
		c.WriteClipboard(strings.Join(recordLines(cur, a.row), "\n"))
	}
}

// rowBadge draws the row number beside the inspector title.
func rowBadge(c *ui.Context, text string) {
	t := c.Theme()
	ui.Row(c).Padding(2, 7).Radius(4).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
		ui.Text(c, text).Font(monoFamily).FontSize(10).TextColor(t.TextMuted).SingleLine()
	})
}

// recordLines formats the row at index as a JSON object, one line per column.
// Numeric columns keep their values unquoted, as JSON numbers.
func recordLines(cur table, index int) []string {
	lines := []string{"{"}
	for i, col := range cur.columns {
		key, _ := json.Marshal(col.name)
		sep := ","
		if i == len(cur.columns)-1 {
			sep = ""
		}
		lines = append(lines, fmt.Sprintf("  %s: %s%s", key, jsonValue(col.typ, cur.rows[index][i]), sep))
	}
	return append(lines, "}")
}

// jsonValue formats one cell value for a column of the given SQL type.
func jsonValue(typ, value string) string {
	if isNumeric(typ) && json.Valid([]byte(value)) {
		return value
	}
	quoted, _ := json.Marshal(value)
	return string(quoted)
}

// isNumeric reports whether a SQL type holds numbers.
func isNumeric(typ string) bool {
	for _, prefix := range []string{"bigint", "int", "numeric", "float", "double", "real"} {
		if strings.HasPrefix(typ, prefix) {
			return true
		}
	}
	return false
}
