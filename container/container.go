package container

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func RunDemo() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	content := container.New(layout.NewGridLayout(2),
		widget.NewButton("Without layout", func() {
			withoutLayoutDemo(myApp)
		}),
		widget.NewButton("With Grid Layout", func() {
			withGridLayoutDemo(myApp)
		}),
		widget.NewButton("With Form Layout", func() {
			withFormLayoutDemo(myApp)
		}),
		widget.NewButton("With HBox Layout", func() {
			withHBoxLayoutDemo(myApp)
		}),
		widget.NewButton("With VBox Layout", func() {
			withVBoxLayoutDemo(myApp)
		}),
	)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func withoutLayoutDemo(a fyne.App) {
	w := a.NewWindow("Without layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text3 := canvas.NewText("World", green)
	text2.Move(fyne.NewPos(20, 20))
	text3.Move(fyne.NewPos(40, 40))
	// container.NewWithoutLayout is a container that does not use layout
	// so we can move text2 to the right of text1
	content := container.NewWithoutLayout(text1, text2, text3)
	w.SetContent(content)
	w.Show()
}

func withGridLayoutDemo(a fyne.App) {
	w := a.NewWindow("With Grid Layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text3 := canvas.NewText("World", green)
	// layout.NewGridLayout(3) is a grid layout with 3 columns
	content := container.New(layout.NewGridLayout(3), text1, text2, text3)
	w.SetContent(content)
	w.Show()
}

func withFormLayoutDemo(a fyne.App) {
	w := a.NewWindow("With Form Layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text3 := canvas.NewText("World", green)
	content := container.New(layout.NewFormLayout(), text1, text2, text3)
	w.SetContent(content)
	w.Show()
}

func withHBoxLayoutDemo(a fyne.App) {
	w := a.NewWindow("With HBox Layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text3 := canvas.NewText("World", green)
	content := container.New(layout.NewHBoxLayout(), text1, text2, text3)
	w.SetContent(content)
	w.Show()
}

func withVBoxLayoutDemo(a fyne.App) {
	w := a.NewWindow("With VBox Layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text3 := canvas.NewText("World", green)
	content := container.New(layout.NewVBoxLayout(), text1, text2, text3)
	w.SetContent(content)
	w.Show()
}
