package basic

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func WindowDemo(a fyne.App) fyne.CanvasObject {
	return widget.NewButton("Window", func() {
		windowDemo(a)
	})
}

func windowDemo(a fyne.App) {
	// create window to show sth.
	win := a.NewWindow("Window")
	// create a label as content of window
	l := widget.NewLabel("This is a fyne demo!")
	// new update button
	b := widget.NewButton("Update", func() {
		tm := time.Now().Format("Time: 03:04:05")
		l.SetText(tm)
	})
	// window size
	win.Resize(fyne.NewSize(200, 100))
	// create a new window top on current window
	openBtn := widget.NewButton("Open Window", func() {
		topWin := a.NewWindow("Top Window")
		topWin.SetContent(widget.NewLabel("This is a top window"))
		topWin.Resize(fyne.NewSize(200, 100))
		topWin.Show()
	})
	// add label and button to window
	win.SetContent(container.NewVBox(l, b, openBtn))
	// show window and run the app
	win.Show()
	// hook method after exited
	tidyup()
}

func tidyup() {
	fmt.Println("exited")
}
