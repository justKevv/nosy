package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

func main() {
	a := newApp()
	mygo.App.WhenReady(func() {
		if err := registerFonts(); err != nil {
			log.Fatal(err)
		}
		mygo.NewWindow(mygo.WindowOptions{
			Title:          "nosy",
			Width:          1400,
			Height:         600,
			MinWidth:       1300,
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
