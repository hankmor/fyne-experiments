package basic

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MakeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello Fyne!")
	in := widget.NewEntry() // entry 用户输入
	// in输入后更改label的内容
	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}

func Helloworld() {
	a := app.New()
	w := a.NewWindow("Hello World!")
	w.SetContent(container.NewVBox(MakeUI())) // vbox是一个垂直布局的容器
	w.ShowAndRun()
}
