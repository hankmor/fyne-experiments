package canvas

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createAndRun(a fyne.App) {
	win := a.NewWindow("Canvas")
	b0 := widget.NewButton("Not Visit Canvas of Window", func() {
		mw := a.NewWindow("Not Visit Canvas of Window")
		mw.Resize(fyne.NewSize(120, 120))
		// widget is CanvasObject, so it can be set to window
		// We don't need to visit canvas of window usually
		mw.SetContent(widget.NewLabel("widget is CanvasObject"))
		mw.Show()
	})
	// create a button to show a new window with a rectangle
	b1 := widget.NewButton("Rectangle Color", func() {
		mw := a.NewWindow("Rectangle Color")
		mw.Resize(fyne.NewSize(120, 120))
		contentRectangle(mw.Canvas())
		mw.Show()
	})
	// create a button to show a new window with a text
	b2 := widget.NewButton("Text", func() {
		mw := a.NewWindow("Text")
		mw.Resize(fyne.NewSize(120, 120))
		contentText(mw.Canvas())
		mw.Show()
	})
	// create a button to show a new window with a circle
	b3 := widget.NewButton("Circle", func() {
		mw := a.NewWindow("Circle")
		mw.Resize(fyne.NewSize(120, 120))
		contentCircle(mw.Canvas())
		mw.Show()
	})
	win.SetContent(container.NewVBox(b0, b1, b2, b3))
	// get canvas from window
	win.Resize(fyne.NewSize(160, 160))
	win.Show()
}

func contentRectangle(c fyne.Canvas) {
	// create a rectangle, color blue
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	// set content to canvas
	c.SetContent(rect)

	// update rectangle color to green after 1 second
	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		// make sure the update is in fyne.DO method, otherwise it will cause race condition
		fyne.Do(func() {
			rect.FillColor = green // update the color of the rectangle
			rect.Refresh()         // refresh the rectangle to show the new color
		})
	}()
}

func contentText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	text.TextSize = 24
	text.Alignment = fyne.TextAlignCenter
	c.SetContent(text)
}

func contentCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}

func RunDemo(a fyne.App) fyne.CanvasObject {
	return widget.NewButton("Canvas", func() {
		createAndRun(a)
	})
}
