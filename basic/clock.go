package basic

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func ClockDemo(a fyne.App) fyne.CanvasObject {
	return widget.NewButton("Clock", func() {
		clockDemo(a)
	})
}

func clockDemo(a fyne.App) {
	win := a.NewWindow("Clock")
	c := color.NRGBA{G: 0xff, A: 200}
	output := canvas.NewText(time.Now().Format(time.DateTime), c) // green color
	output.TextStyle.Monospace = true                             // monospace font style
	output.TextSize = 32                                          // font size
	win.SetContent(output)

	// update time in goroutine
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			// fyne.Do(func() { // must run in fyne.Do
			fyne.DoAndWait(func() { // 更新完成之前不会返回，可以避免并发时的资源共享带来的问题
				output.Text = time.Now().Format(time.DateTime)
				output.Refresh() // redraw text
			})
		}
	}()
	win.Show()
}
