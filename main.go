package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

// table is one table of the database the window shows. The rows are sample
// data until a connection is implemented.
type table struct {
	name    string
	columns []string
	rows    [][]string
}

// app is the state the window shows. Its view builds the interface from
// it, on the main thread, whenever the window needs a frame: after input,
// after Window.Update, and while something animates.
type app struct {
	tables   []table
	selected string
	search   string
	rows     ui.ListState
	row      int
}

func newApp() *app {
	a := &app{tables: sampleTables(), selected: "users", row: -1}
	a.rows.Selected = &a.row
	return a
}

func (a *app) current() table {
	for _, t := range a.tables {
		if t.name == a.selected {
			return t
		}
	}
	return table{}
}

func (a *app) view(c *ui.Context) {
	t := c.Theme()
	cur := a.current()
	ui.Column(c).Fill().Children(func() {
		a.topBar(c)
		ui.Row(c).Grow(1).AlignItems(ui.Stretch).Children(func() {
			ui.Sidebar(c, &a.selected, func() {
				ui.SidebarSection(c, "Tables", nil, func() {
					for _, tb := range a.tables {
						ui.SidebarItem(c, tb.name, nil, tb.name)
					}
				})
			}).Width(200).OnChange(func() { a.row = -1 })
			ui.Column(c).Grow(1).Children(func() {
				ui.Toolbar(c, func() {
					ui.Text(c, cur.name).FontSize(16).Bold()
					ui.Spacer(c)
					ui.Text(c, "Not connected").TextColor(t.TextMuted)
				})
				columns := make([]ui.TableColumn, len(cur.columns))
				for i, name := range cur.columns {
					columns[i] = ui.TableColumn{Title: name}
				}
				ui.Table(c, &a.rows, columns, len(cur.rows), func(row, col int) {
					ui.Text(c, cur.rows[row][col]).SingleLine()
				}).Grow(1)
				ui.Textf(c, "%d rows · %d columns", len(cur.rows), len(cur.columns)).
					TextColor(t.TextMuted).Padding(6, 12)
			})
		})
	})
}

// lineIcon wraps stroke-drawn icon paths in the stroke style the icons share.
func lineIcon(body string) *ui.SVG {
	return ui.MustParseSVG([]byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">` + body + `</g></svg>`))
}

var (
	databaseIcon = lineIcon(`<ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5v14a9 3 0 0 0 18 0V5"/><path d="M3 12a9 3 0 0 0 18 0"/>`)
	chevronIcon  = lineIcon(`<path d="m6 9 6 6 6-6"/>`)
	activityIcon = lineIcon(`<path d="M22 12h-4l-3 9L9 3l-3 9H2"/>`)
)

// titleBarHeight is the height of the top bar. It holds the native window
// controls, which are centered in it.
const titleBarHeight = 44

// topBar draws the window's top row in the title bar, beside the window
// controls, which it leaves room for.
func (a *app) topBar(c *ui.Context) {
	t := c.Theme()
	bar := c.TitleBar()
	ui.Row(c).Height(titleBarHeight).Padding(0, bar.Right+15, 0, bar.Left+20).Gap(10).AlignItems(ui.Center).DragWindow().Children(func() {
		ui.Row(c).Grow(1).Gap(10).AlignItems(ui.Center).Children(func() {
			ui.Row(c).Size(24, 24).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
				ui.Icon(c, databaseIcon).Size(14, 14).TextColor(t.Text)
			})
			ui.Text(c, "nosy").Bold().TextColor(t.TextMuted)
			ui.Row(c).Padding(6, 10).Gap(6).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
				ui.Text(c, "● prod-us-east-1 / telemetry_core").SingleLine()
				ui.Text(c, "PG 16.4").FontSize(11).TextColor(t.TextMuted)
				ui.Icon(c, chevronIcon).Size(12, 12).TextColor(t.TextMuted)
			})
		})
		ui.Row(c).Grow(1).Gap(10).AlignItems(ui.Center).Children(func() {
			ui.Spacer(c)
			ui.TextInput(c, &a.search).Placeholder("Jump to table, index, or SQL…").Width(300)
			ui.Row(c).Padding(6, 10).Gap(6).Radius(6).Border(1, t.Border).Background(t.Surface).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, activityIcon).Size(12, 12).TextColor(t.TextMuted)
				ui.Text(c, "14/50 conns · 1.8ms").TextColor(t.TextMuted).SingleLine()
			})
			ui.PrimaryButton(c, "+ New Query")
		})
	})
}

func sampleTables() []table {
	return []table{
		{
			name:    "users",
			columns: []string{"id", "name", "email", "created_at"},
			rows: [][]string{
				{"1", "Ada Lovelace", "ada@example.com", "2026-01-04"},
				{"2", "Alan Turing", "alan@example.com", "2026-01-09"},
				{"3", "Grace Hopper", "grace@example.com", "2026-02-17"},
				{"4", "Linus Torvalds", "linus@example.com", "2026-03-02"},
				{"5", "Margaret Hamilton", "margaret@example.com", "2026-03-21"},
			},
		},
		{
			name:    "orders",
			columns: []string{"id", "user_id", "total", "status"},
			rows: [][]string{
				{"1001", "1", "42.50", "shipped"},
				{"1002", "3", "13.00", "pending"},
				{"1003", "2", "199.99", "delivered"},
				{"1004", "5", "7.25", "cancelled"},
			},
		},
		{
			name:    "products",
			columns: []string{"id", "sku", "name", "price"},
			rows: [][]string{
				{"1", "KB-100", "Mechanical keyboard", "89.00"},
				{"2", "MS-220", "Wireless mouse", "29.90"},
				{"3", "MN-027", "27-inch monitor", "249.00"},
			},
		},
	}
}

func main() {
	a := newApp()
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:          "nosy",
			Width:          960,
			Height:         600,
			MinWidth:       640,
			MinHeight:      400,
			TitleBarStyle:  mygo.TitleBarHidden,
			TitleBarHeight: titleBarHeight,
			// Centers the macOS traffic lights (about 14px tall) in the bar.
			TrafficLightPosition: &mygo.Point{X: 16, Y: 15},
			// Opens where the user left it last time.
			StateKey: "main",
			// The window shows the interface MyGo draws, not a web page.
			Content: ui.View(a.view),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
