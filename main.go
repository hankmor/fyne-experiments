package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// create app
	ap := app.New()
	// create window to show sth.
	w := ap.NewWindow("Hello fyne")
	// create a label as content of window
	l := widget.NewLabel("This is a fyne demo!")
	// new update button
	b := widget.NewButton("Update", func() {
		tm := time.Now().Format("Time: 03:04:05")
		l.SetText(tm)
	})
	// add label and button to window
	w.SetContent(container.NewVBox(l, b))
	// window size
	w.Resize(fyne.NewSize(200, 100))
	// show window and run the app
	w.ShowAndRun()
	// hook method after exited
	tidyup()
}

func tidyup() {
	fmt.Println("exited")
}
