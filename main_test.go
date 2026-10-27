package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// The view runs without a window in tests, which click and type as a user
// would.
func TestView(t *testing.T) {
	a := newApp()
	tt := ui.NewTester(a.view, 960, 600)
	if !tt.HasText("Ada Lovelace") {
		t.Fatalf("users rows missing, texts %q", tt.Texts())
	}
	if err := tt.Click("orders"); err != nil {
		t.Fatal(err)
	}
	if a.selected != "orders" || !tt.HasText("shipped") || tt.HasText("Ada Lovelace") {
		t.Errorf("selected %q after choosing orders, texts %q", a.selected, tt.Texts())
	}
}

func TestWorkspace(t *testing.T) {
	a := newApp()
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
