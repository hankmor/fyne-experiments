package basic

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func ClockDemo() {
	ap := app.New()
	w := ap.NewWindow("Clock")

	c := color.NRGBA{G: 0xff, A: 200}
	output := canvas.NewText(time.Now().Format(time.DateTime), c) // green color
	output.TextStyle.Monospace = true                             // monospace font style
	output.TextSize = 32                                          // font size
	w.SetContent(output)

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
	w.ShowAndRun()
}
