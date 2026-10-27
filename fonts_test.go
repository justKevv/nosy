package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func TestRegisterFonts(t *testing.T) {
	if err := registerFonts(); err != nil {
		t.Fatal(err)
	}
	for _, f := range ui.FontFamilies() {
		if f.Name == monoFamily {
			if !f.Registered || !f.Monospaced {
				t.Errorf("%s registered %v, monospaced %v", monoFamily, f.Registered, f.Monospaced)
			}
			return
		}
	}
	t.Fatalf("%s not in the font families", monoFamily)
}
