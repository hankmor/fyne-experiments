package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ThemeDemo(a fyne.App) fyne.CanvasObject {
	return widget.NewButton("Theme", func() {
		customThemeDemo(a)
	})
}
