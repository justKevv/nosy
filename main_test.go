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
