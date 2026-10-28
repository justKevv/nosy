package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// The view runs without a window in tests, which click and type as a user
// would.
func TestView(t *testing.T) {
	a := newApp()
	a.screen = screenWorkspace
	tt := ui.NewTester(a.view, 960, 600)
	if !tt.HasText("Ada Lovelace") {
		t.Fatalf("users rows missing, texts %q", tt.Texts())
	}
	if err := tt.Click("orders"); err != nil {
		t.Fatal(err)
	}
	if a.selected != "orders" || !tt.HasText("SHIPPED") || tt.HasText("Ada Lovelace") {
		t.Errorf("selected %q after choosing orders, texts %q", a.selected, tt.Texts())
	}
}

func TestStart(t *testing.T) {
	a := newApp()
	tt := ui.NewTester(a.view, 1440, 900)
	for _, text := range []string{"New connection", "DRIVER", "Require SSL", "QUICK RECONNECT", "PostgreSQL 16.4"} {
		if !tt.HasText(text) {
			t.Fatalf("%q missing, texts %q", text, tt.Texts())
		}
	}
	if err := tt.Click("staging-use1"); err != nil {
		t.Fatal(err)
	}
	if a.picked != 2 || a.host != "staging-use1.db.internal" || a.drivers[a.driver].name != "MongoDB" {
		t.Errorf("staging-use1 not loaded: picked %d, host %q, driver %d", a.picked, a.host, a.driver)
	}
	if err := tt.Click("Connect"); err != nil {
		t.Fatal(err)
	}
	if a.screen != screenWorkspace {
		t.Errorf("screen %d after connecting, want the workspace", a.screen)
	}
	if a.target == nil || a.target.name != "staging-use1" || a.target.database != "sessions" {
		t.Errorf("target %+v after connecting staging-use1", a.target)
	}
}

func TestConnectFieldFocus(t *testing.T) {
	a := newApp()
	tt := ui.NewTester(a.view, 1440, 900)
	label, ok := tt.Find("HOST")
	if !ok {
		t.Fatalf("HOST label missing, texts %q", tt.Texts())
	}
	tt.ClickAt(label.X+label.W/2, label.Y+label.H+6+20)
	if active, _ := tt.TextInputActive(); !active {
		t.Fatal("host field not focused after click")
	}
	tt.Type("db")
	if a.host != "db" {
		t.Errorf("host %q after typing, want db", a.host)
	}
}

func TestWorkspace(t *testing.T) {
	a := newApp()
	a.screen = screenWorkspace
	tt := ui.NewTester(a.view, 1400, 800)
	for _, text := range []string{"Run Selection ⌘↩", "WHERE", "Showing 1–5 of 5 rows", "ROW INSPECTOR", "Copy JSON", "Select a row to inspect its values."} {
		if !tt.HasText(text) {
			t.Fatalf("%q missing, texts %q", text, tt.Texts())
		}
	}
	if err := tt.Click("Ada Lovelace"); err != nil {
		t.Fatal(err)
	}
	if a.row < 0 || tt.HasText("Select a row to inspect its values.") || !tt.HasText("ada@example.com") {
		t.Errorf("row %d not inspected, texts %q", a.row, tt.Texts())
	}
	if !tt.HasText("Row #1") || !tt.HasText("COLUMNS & TYPES") || !tt.HasText("RECORD JSON") || !tt.HasText("\"id\": 1,") {
		t.Errorf("record JSON missing, texts %q", tt.Texts())
	}
}
