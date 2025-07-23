package theme

import (
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct {
	fyne.Theme
}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameHyperlink {
		return color.RGBA{R: 255, G: 0, B: 0, A: 255} // 红色超链接
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func customThemeDemo(a fyne.App) {
	a.Settings().SetTheme(myTheme{}) // 设置自定义主题, 这是全局设置，所有的超链接都会使用这个主题
	win := a.NewWindow("Custom Theme Hyperlink")
	win.Resize(fyne.NewSize(400, 200))

	// 创建 URL
	url, err := url.Parse("https://hankmo.com")
	if err != nil {
		fyne.LogError("Failed to parse URL", err)
		return
	}

	// 创建超链接
	hyperlink := widget.NewHyperlinkWithStyle(
		"https://hankmo.com",
		url,
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	// 设置窗口内容
	win.SetContent(container.NewVBox(
		hyperlink,
		widget.NewLabel("Red hyperlink with custom theme"),
	))

	win.Show()
}
