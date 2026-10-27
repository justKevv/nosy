package main

import (
	"embed"
	"fmt"

	"github.com/egoist/mygo/ui"
)

// monoFamily is the family of the monospaced text of the app.
const monoFamily = "JetBrains Mono"

// monoFonts are the weights of monoFamily the app bundles, in fonts/.
var monoFonts = []string{
	"fonts/JetBrainsMono-Regular.ttf",
	"fonts/JetBrainsMono-Medium.ttf",
	"fonts/JetBrainsMono-SemiBold.ttf",
	"fonts/JetBrainsMono-Bold.ttf",
}

//go:embed fonts/*.ttf
var fontFiles embed.FS

// registerFonts makes monoFamily available to text. The fonts are bundled
// because the app cannot count on them being installed.
func registerFonts() error {
	for _, name := range monoFonts {
		data, err := fontFiles.ReadFile(name)
		if err != nil {
			return fmt.Errorf("read %s: %w", name, err)
		}
		if err := ui.RegisterFont(data, monoFamily); err != nil {
			return fmt.Errorf("register %s: %w", name, err)
		}
	}
	return nil
}
