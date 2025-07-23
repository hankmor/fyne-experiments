package widget

import (
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// 自定义超链接, 扩展Hyperlink, 添加背景颜色
type CustomHyperlink struct {
	widget.Hyperlink
	BackgroundColor color.Color // 背景颜色
}

// 自定义渲染器
type customHyperlinkRenderer struct {
	hyperlink *CustomHyperlink
	text      *canvas.Text
	objects   []fyne.CanvasObject
}

func (r *customHyperlinkRenderer) Destroy() {}

func (r *customHyperlinkRenderer) Layout(size fyne.Size) {
	r.text.Resize(size)
}

func (r *customHyperlinkRenderer) MinSize() fyne.Size {
	return r.text.MinSize()
}

func (r *customHyperlinkRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *customHyperlinkRenderer) Refresh() {
	r.text.Text = r.hyperlink.Text
	r.text.Alignment = r.hyperlink.Alignment
	r.text.TextStyle = r.hyperlink.TextStyle
	r.text.Color = color.RGBA{R: 0, G: 128, B: 0, A: 255} // 自定义文本颜色
	// 不设置下划线和鼠标样式
	canvas.Refresh(r.hyperlink)
}

func (r *customHyperlinkRenderer) BackgroundColor() color.Color {
	return r.hyperlink.BackgroundColor
}

func (r *customHyperlinkRenderer) ApplyTheme() {
	// 自定义主题应用（可选）
}

// 创建自定义渲染器
func (h *CustomHyperlink) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(h.Text, color.RGBA{R: 0, G: 128, B: 0, A: 255})
	text.Alignment = h.Alignment
	text.TextStyle = h.TextStyle
	return &customHyperlinkRenderer{
		hyperlink: h,
		text:      text,
		objects:   []fyne.CanvasObject{text},
	}
}

func NewCustomHyperlink(text string, url *url.URL) *CustomHyperlink {
	h := &CustomHyperlink{
		Hyperlink:       *widget.NewHyperlink(text, url),
		BackgroundColor: color.RGBA{R: 200, G: 200, B: 200, A: 255}, // 灰色背景
	}
	h.ExtendBaseWidget(h)
	return h
}
