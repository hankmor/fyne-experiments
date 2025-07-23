package basic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello Fyne!")
	in := widget.NewEntry() // entry 用户输入
	// in输入后更改label的内容
	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}

func Helloworld(a fyne.App) fyne.CanvasObject {
	return widget.NewButton("HelloWorld", func() {
		win := a.NewWindow("HelloWorld")
		win.SetContent(container.NewVBox(makeUI()))
		win.Show()
	})
}
