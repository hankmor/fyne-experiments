package main

import (
	"fyne-experiments/basic"
	"fyne-experiments/canvas"
	"fyne-experiments/container"
	"fyne-experiments/theme"
	"fyne-experiments/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	c "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func createApp() (fyne.App, fyne.Window) {
	app := app.New()
	// app.Settings().SetTheme(theme.LightTheme())
	window := app.NewWindow("Fyne Experiments")
	return app, window
}

func main() {
	a, w := createApp()
	c := c.New(layout.NewGridLayout(2))

	c.Add(basic.Helloworld(a))
	c.Add(basic.WindowDemo(a))
	c.Add(basic.ClockDemo(a))
	c.Add(canvas.RunDemo(a))
	c.Add(container.RunDemo(a))
	c.Add(widget.RunDemo(a))
	c.Add(theme.ThemeDemo(a))

	w.SetContent(c)
	w.ShowAndRun()
}
