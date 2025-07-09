package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// create app
	ap := app.New()
	// create window to show sth.
	w := ap.NewWindow("Hello fyne")
	// create a label as content of window
	w.SetContent(widget.NewLabel("This is a fyne demo!"))
	// show window and run the app
	w.ShowAndRun()
}
