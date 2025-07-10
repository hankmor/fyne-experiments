package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func windowDemo() {
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
	// window size
	w.Resize(fyne.NewSize(200, 100))
	// create a new window top on current window
	openBtn := widget.NewButton("Open Window", func() {
		w := ap.NewWindow("Top Window")
		w.SetContent(widget.NewLabel("This is a top window"))
		w.Resize(fyne.NewSize(200, 100))
		w.Show()
	})
	// add label and button to window
	w.SetContent(container.NewVBox(l, b, openBtn))
	// show window and run the app
	w.ShowAndRun()
	// hook method after exited
	tidyup()
}

func tidyup() {
	fmt.Println("exited")
}
