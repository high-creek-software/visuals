package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/high-creek-software/visuals/internal"
)

func main() {
	a := app.NewWithID("io.highcreeksoftware.visuals")
	win := a.NewWindow("Visuals")
	win.Resize(fyne.NewSize(1200, 550))

	v := internal.NewVisualsApp(a, win)
	v.Start()
}
