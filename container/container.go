package container

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func RunDemo(a fyne.App) fyne.CanvasObject {
	return widget.NewButton("Container", func() {
		containerDemo(a)
	})
}

func containerDemo(a fyne.App) {
	win := a.NewWindow("Container")
	content := container.New(layout.NewGridLayout(2),
		widget.NewButton("Without layout", func() {
			withoutLayoutDemo(a)
		}),
		widget.NewButton("With Grid Layout", func() {
			withGridLayoutDemo(a)
		}),
		widget.NewButton("With Form Layout", func() {
			withFormLayoutDemo(a)
		}),
		widget.NewButton("With HBox Layout", func() {
			withHBoxLayoutDemo(a)
		}),
		widget.NewButton("With VBox Layout", func() {
			withVBoxLayoutDemo(a)
		}),
	)
	win.SetContent(content)
	win.Show()
}

func withoutLayoutDemo(a fyne.App) {
	win := a.NewWindow("Without layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text3 := canvas.NewText("World", green)
	text2.Move(fyne.NewPos(20, 20))
	text3.Move(fyne.NewPos(40, 40))
	// container.NewWithoutLayout is a container that does not use layout
	// so we can move text2 to the right of text1
	content := container.NewWithoutLayout(text1, text2, text3)
	win.SetContent(content)
	win.Show()
}

func withGridLayoutDemo(a fyne.App) {
	win := a.NewWindow("With Grid Layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text3 := canvas.NewText("World", green)
	// layout.NewGridLayout(3) is a grid layout with 3 columns
	content := container.New(layout.NewGridLayout(3), text1, text2, text3)
	win.SetContent(content)
	win.Show()
}

func withFormLayoutDemo(a fyne.App) {
	win := a.NewWindow("With Form Layout")
	text1 := widget.NewLabel("Name")
	entry1 := widget.NewEntry()
	text2 := widget.NewLabel("Age")
	entry2 := widget.NewEntry()
	text3 := widget.NewLabel("Gender")
	radio1 := widget.NewRadioGroup([]string{"Male", "Female"}, func(s string) {
		println("Gender:", s)
	})
	content := container.New(layout.NewFormLayout(), text1, entry1, text2, entry2, text3, radio1)
	win.SetContent(content)
	win.Show()
}

func withHBoxLayoutDemo(a fyne.App) {
	win := a.NewWindow("With HBox Layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text3 := canvas.NewText("World", green)
	content := container.New(layout.NewHBoxLayout(), text1, text2, text3)
	win.SetContent(content)
	win.Show()
}

func withVBoxLayoutDemo(a fyne.App) {
	win := a.NewWindow("With VBox Layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text3 := canvas.NewText("World", green)
	content := container.New(layout.NewVBoxLayout(), text1, text2, text3)
	win.SetContent(content)
	win.Show()
}
