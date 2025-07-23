package widget

import (
	"io"
	"net/url"
	"os"
	"time"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func RunDemo(a fyne.App) fyne.CanvasObject {
	return widget.NewButton("Widget", func() {
		widgetDemo(a)
	})
}

func widgetDemo(a fyne.App) {
	win := a.NewWindow("Widget")
	win.SetContent(container.New(layout.NewGridLayout(4),
		widget.NewButton("Accordion", func() {
			accordion(a)
		}),
		widget.NewButton("Activity", func() {
			activity(a)
		}),
		widget.NewButton("Button", func() {
			button(a)
		}),
		widget.NewButton("Card", func() {
			card(a)
		}),
		widget.NewButton("CheckBox", func() {
			checkBox(a)
		}),
		widget.NewButton("Entry", func() {
			entry(a)
		}),
		widget.NewButton("FileIcon", func() {
			fileicon(a)
		}),
		widget.NewButton("Form", func() {
			form(a)
		}),
		widget.NewButton("Hyperlink", func() {
			hyperlink(a)
		}),
		widget.NewButton("Icon", func() {
			icon(a)
		}),
		widget.NewButton("Label", func() {
			lable(a)
		}),
		widget.NewButton("ProgressBar", func() {
			progressBar(a)
		}),
		widget.NewButton("RadioGroup", func() {
			radioGroup(a)
		}),
		widget.NewButton("RichText", func() {
			richText(a)
		}),
		widget.NewButton("Select", func() {
			selectWidget(a)
		}),
		widget.NewButton("SelectEntry", func() {
			selectEntry(a)
		}),
		widget.NewButton("Separator", func() {
			separator(a)
		}),
	))
	win.Show()
}

// accordion is a widget that displays a list of items in a collapsible section
func accordion(a fyne.App) {
	w := a.NewWindow("Accordion")
	w.SetContent(widget.NewAccordion(
		// add a famous quote
		widget.NewAccordionItem("In Life", widget.NewLabel("In life, as in art, the beautiful moves in curves.")),
		widget.NewAccordionItem("In Art", widget.NewLabel("Of all possessions a friend is the most precious.")),
		widget.NewAccordionItem("In Friendship", widget.NewLabel("A true friend is someone who is there for you when he'd rather be anywhere else.")),
		widget.NewAccordionItem("In Love", widget.NewLabel("Love is the only force capable of transforming an enemy into a friend.")),
		widget.NewAccordionItem("In Music", widget.NewLabel("Music is the universal language of mankind.")),
		widget.NewAccordionItem("In Literature", widget.NewLabel("The only way to do great work is to love what you do.")),
	))
	w.Show()
}

// activity is a widget that displays a spinning activity indicator
func activity(a fyne.App) {
	w := a.NewWindow("Activity")
	activity := widget.NewActivity()
	activity.Start()
	w.SetContent(activity)
	go func() {
		time.Sleep(3 * time.Second)
		activity.Stop()
	}()
	w.Show()
}

func button(a fyne.App) {
	w := a.NewWindow("Button")
	// normal button
	normalButton := widget.NewButton("Button", func() {
		println("Button clicked")
	})
	// button with icon
	icon := theme.CancelIcon()
	button := widget.NewButtonWithIcon("Cancel", icon, func() {
		println("Cancel button clicked")
	})
	w.SetContent(container.NewVBox(normalButton, button))
	w.Show()
}

func card(a fyne.App) {
	w := a.NewWindow("Card")

	c1 := widget.NewCard("In Life", "Life is art", widget.NewLabel("In life, as in art, the beautiful moves in curves."))

	// bf.jpg is a 40x40 image
	f, err := os.Open("bf.jpg")
	if err != nil {
		println("Error opening image:", err)
		return
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		println("Error reading image:", err)
		return
	}
	img := canvas.NewImageFromResource(fyne.NewStaticResource("image", bytes))
	// set the image size
	// img.Resize(fyne.NewSize(40, 40)) // fix the image size
	img.SetMinSize(fyne.NewSize(0, 120))   // set the minimum size, width is 0, so the image will auto adjust the width
	img.FillMode = canvas.ImageFillContain // keep the aspect ratio, fully display the image
	// or use other fill modes:
	// img.FillMode = canvas.ImageFillStretch // stretch to fill the entire area
	// img.FillMode = canvas.ImageFillOriginal // keep the original size
	// put the image at the top
	// create a custom card layout: image at the top, then title and subtitle
	title := canvas.NewText("In Art", nil)
	title.TextSize = 18                          // set the title font size
	title.TextStyle = fyne.TextStyle{Bold: true} // set the title to bold

	subtitle := canvas.NewText("Art is life", nil)
	subtitle.TextSize = 14                            // set the subtitle font size
	subtitle.TextStyle = fyne.TextStyle{Italic: true} // set the subtitle to italic

	c2Content := container.NewVBox(
		img,
		title,
		subtitle,
	)
	c2 := widget.NewCard("", "", nil)
	c2.SetContent(c2Content)

	// 让窗口自动适应内容大小，或者设置合适的固定大小
	// 或者不设置大小，让窗口自动调整
	w.SetContent(container.New(layout.NewVBoxLayout(), c1, c2))
	w.Show()
}

func checkBox(a fyne.App) {
	w := a.NewWindow("CheckBox")
	checkBox := widget.NewCheck("I agree to the terms and conditions", func(b bool) {
		println("CheckBox value:", b)
	})
	checkBox.SetChecked(true) // default checked
	w.SetContent(container.NewVBox(checkBox))
	w.Show()
}

func entry(a fyne.App) {
	w := a.NewWindow("Entry")
	entry := widget.NewEntry()
	entry.SetText("Hello, World!")
	entry.OnChanged = func(text string) {
		println("Entry value:", text)
	}
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetText("123456")
	passwordEntry.OnChanged = func(text string) {
		println("PasswordEntry value:", text)
	}
	selectEntry := widget.NewSelectEntry([]string{"Hello", "World", "Fyne"})
	selectEntry.OnChanged = func(text string) {
		println("SelectEntry value:", text)
	}
	w.SetContent(container.NewVBox(entry, passwordEntry, selectEntry))
	w.Resize(fyne.NewSize(200, 100))
	w.Show()
}

func fileicon(a fyne.App) {
	w := a.NewWindow("FileIcon")

	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		println("Error getting current directory:", err)
		return
	}
	println("Current directory:", currentDir)

	// 方法1: 使用系统文件图标（根据文件扩展名自动选择）
	systemFileIcon := widget.NewFileIcon(nil)
	systemFileIcon.SetURI(storage.NewFileURI("go.mod")) // 项目根目录的go.mod文件

	systemFolderIcon := widget.NewFileIcon(nil)
	systemFolderIcon.SetURI(storage.NewFileURI("widget")) // 项目下的widget文件夹

	systemImageIcon := widget.NewFileIcon(nil)
	systemImageIcon.SetURI(storage.NewFileURI("bf.jpg")) // 项目根目录的图片文件

	// 方法2: 使用Fyne主题图标（系统风格但统一）
	themeFileIcon := widget.NewButtonWithIcon("", theme.DocumentIcon(), func() {})
	themeFileIcon.Importance = widget.LowImportance // 移除按钮样式

	themeFolderIcon := widget.NewButtonWithIcon("", theme.FolderIcon(), func() {})
	themeFolderIcon.Importance = widget.LowImportance

	themeImageIcon := widget.NewButtonWithIcon("", theme.ComputerIcon(), func() {})
	themeImageIcon.Importance = widget.LowImportance

	// 方法3: 使用canvas.Image显示主题图标
	canvasFileIcon := canvas.NewImageFromResource(theme.DocumentIcon())
	canvasFileIcon.SetMinSize(fyne.NewSize(32, 32))
	canvasFileIcon.FillMode = canvas.ImageFillContain

	canvasFolderIcon := canvas.NewImageFromResource(theme.FolderIcon())
	canvasFolderIcon.SetMinSize(fyne.NewSize(32, 32))
	canvasFolderIcon.FillMode = canvas.ImageFillContain

	canvasImageIcon := canvas.NewImageFromResource(theme.ComputerIcon())
	canvasImageIcon.SetMinSize(fyne.NewSize(32, 32))
	canvasImageIcon.FillMode = canvas.ImageFillContain

	// 方法4: 使用不同大小的系统图标
	largeSystemIcon := widget.NewFileIcon(nil)
	largeSystemIcon.SetURI(storage.NewFileURI("go.mod"))
	largeSystemIcon.Resize(fyne.NewSize(48, 48))

	smallSystemIcon := widget.NewFileIcon(nil)
	smallSystemIcon.SetURI(storage.NewFileURI("main.go"))
	smallSystemIcon.Resize(fyne.NewSize(16, 16))

	// 方法5: 动态获取项目目录下的文件和文件夹
	var projectIcons []fyne.CanvasObject

	// 读取当前目录下的文件和文件夹
	entries, err := os.ReadDir(".")
	if err != nil {
		println("Error reading directory:", err)
	} else {
		// 只显示前6个项目
		count := 0
		for _, entry := range entries {
			if count >= 6 {
				break
			}

			// 创建文件图标
			icon := widget.NewFileIcon(nil)
			icon.SetURI(storage.NewFileURI(entry.Name()))

			// 创建标签显示文件名
			label := widget.NewLabel(entry.Name())
			if len(entry.Name()) > 10 {
				label.SetText(entry.Name()[:10] + "...")
			}

			// 添加到容器
			projectIcons = append(projectIcons, container.NewVBox(icon, label))
			count++
		}
	}

	// 创建所有图标容器
	var allIcons []fyne.CanvasObject

	// 添加固定的示例图标
	allIcons = append(allIcons,
		container.NewVBox(systemFileIcon, widget.NewLabel("go.mod")),
		container.NewVBox(systemFolderIcon, widget.NewLabel("widget/")),
		container.NewVBox(systemImageIcon, widget.NewLabel("bf.jpg")),
		container.NewVBox(themeFileIcon, widget.NewLabel("Theme File")),
		container.NewVBox(themeFolderIcon, widget.NewLabel("Theme Folder")),
		container.NewVBox(themeImageIcon, widget.NewLabel("Theme Image")),
		container.NewVBox(canvasFileIcon, widget.NewLabel("Canvas File")),
		container.NewVBox(canvasFolderIcon, widget.NewLabel("Canvas Folder")),
		container.NewVBox(canvasImageIcon, widget.NewLabel("Canvas Image")),
		container.NewVBox(largeSystemIcon, widget.NewLabel("Large go.mod")),
		container.NewVBox(smallSystemIcon, widget.NewLabel("Small main.go")),
	)

	// 添加动态获取的项目文件图标
	allIcons = append(allIcons, projectIcons...)

	// 使用网格布局排列所有图标
	content := container.NewGridWithColumns(4, allIcons...)

	w.SetContent(content)
	w.Resize(fyne.NewSize(500, 400))
	w.Show()
}

func form(a fyne.App) {
	w := a.NewWindow("Form")
	form := widget.NewForm(
		widget.NewFormItem("Name", widget.NewEntry()),
		widget.NewFormItem("Email", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)
	w.SetContent(form)
	w.Resize(fyne.NewSize(200, 100))
	w.Show()
}

func hyperlink(a fyne.App) {
	w := a.NewWindow("Hyperlink")

	// 方法1: 使用widget.NewHyperlink (基础方法，样式选项有限)
	basicHyperlink := widget.NewHyperlink("https://hankmo.com", &url.URL{
		Scheme: "https",
		Host:   "hankmo.com",
	})
	basicHyperlink.Alignment = fyne.TextAlignCenter
	// 设置超链接字体样式 (可用的样式选项)
	basicHyperlink.TextStyle = fyne.TextStyle{
		Bold:      true,
		Monospace: true,
		Symbol:    true,
		Underline: true,
	}

	// 方法2: 使用canvas.Text创建自定义样式的超链接 (推荐方法)
	// 创建带样式的文本
	blueColor := color.NRGBA{R: 0, G: 0, B: 255, A: 255} // 蓝色
	customText := canvas.NewText("https://hankmo.com", blueColor)
	customText.TextSize = 18 // 设置字号
	customText.TextStyle = fyne.TextStyle{
		Bold:      true,
		Underline: true,
	}
	customText.Alignment = fyne.TextAlignCenter

	// 方法3: 使用不同颜色和大小的超链接
	redColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255} // 红色
	largeHyperlink := canvas.NewText("大号超链接", redColor)
	largeHyperlink.TextSize = 24
	largeHyperlink.TextStyle = fyne.TextStyle{Bold: true, Underline: true}
	largeHyperlink.Alignment = fyne.TextAlignCenter

	greenColor := color.NRGBA{R: 0, G: 255, B: 0, A: 255} // 绿色
	smallHyperlink := canvas.NewText("小号超链接", greenColor)
	smallHyperlink.TextSize = 12
	smallHyperlink.TextStyle = fyne.TextStyle{Italic: true, Underline: true}
	smallHyperlink.Alignment = fyne.TextAlignCenter

	// 方法4: 使用自定义颜色
	purpleColor := color.NRGBA{R: 128, G: 0, B: 128, A: 255} // 紫色
	customColorHyperlink := canvas.NewText("自定义颜色超链接", purpleColor)
	customColorHyperlink.TextSize = 16
	customColorHyperlink.TextStyle = fyne.TextStyle{Bold: true, Underline: true}
	customColorHyperlink.Alignment = fyne.TextAlignCenter

	// 方法5: 使用按钮样式的超链接
	buttonStyleHyperlink := widget.NewButton("按钮样式超链接", func() {
		u, _ := url.Parse("https://hankmo.com")
		fyne.CurrentApp().OpenURL(u)
	})
	buttonStyleHyperlink.Importance = widget.LowImportance // 移除按钮边框

	// 方法6: 使用自定义超链接
	customHyperlink := NewCustomHyperlink("自定义超链接", &url.URL{
		Scheme: "https",
		Host:   "hankmo.com",
	})

	// 将所有超链接放在容器中
	content := container.NewVBox(
		widget.NewLabel("基础超链接(样式有限):"), basicHyperlink,
		widget.NewLabel("自定义样式文本 (仅显示):"), customText,
		widget.NewLabel("大号文本 (仅显示):"), largeHyperlink,
		widget.NewLabel("小号文本 (仅显示):"), smallHyperlink,
		widget.NewLabel("自定义颜色文本 (仅显示):"), customColorHyperlink,
		widget.NewLabel("按钮样式超链接:"), buttonStyleHyperlink,
		widget.NewLabel("自定义超链接控件:"), customHyperlink,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 500))
	w.Show()
}

func icon(a fyne.App) {
	w := a.NewWindow("Icon")
	cancelIcon := widget.NewIcon(theme.CancelIcon())
	cancelContainer := container.NewVBox(
		cancelIcon,
		widget.NewLabel("Cancel Icon"),
	)
	okIcon := widget.NewIcon(theme.ConfirmIcon())
	okContainer := container.NewVBox(
		okIcon,
		widget.NewLabel("OK Icon"),
	)
	w.SetContent(container.NewHBox(cancelContainer, okContainer))
	w.Resize(fyne.NewSize(200, 100))
	w.Show()
}

func lable(a fyne.App) {
	w := a.NewWindow("Label")
	w.SetContent(widget.NewLabel("Hello, World!"))
	w.Resize(fyne.NewSize(200, 100))
	w.Show()
}

func progressBar(a fyne.App) {
	w := a.NewWindow("ProgressBar")
	progressBar := widget.NewProgressBar()
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			progressBar.SetValue(float64(i+1) / 10)
			// progressBar.Refresh() // no need to refresh, it will be refreshed after SetValue automatically
		}
	}()
	progressBarInfinite := widget.NewProgressBarInfinite()
	progressBarInfinite.Start()
	go func() {
		time.Sleep(10 * time.Second)
		progressBarInfinite.Stop()
	}()
	w.SetContent(container.NewVBox(progressBar, progressBarInfinite))
	w.Resize(fyne.NewSize(200, 100))
	w.Show()
}

func radioGroup(a fyne.App) {
	w := a.NewWindow("RadioGroup")
	radioGroup := widget.NewRadioGroup([]string{"Option 1", "Option 2", "Option 3"}, func(s string) {
		println("RadioGroup value:", s)
	})
	w.SetContent(radioGroup)
	w.Resize(fyne.NewSize(200, 100))
	w.Show()
}

func richText(a fyne.App) {
	w := a.NewWindow("RichText")
	richText := widget.NewRichText()
	richText.Segments = append(richText.Segments, &widget.TextSegment{
		Style: widget.RichTextStyle{
			Alignment: fyne.TextAlignCenter,
			ColorName: theme.ColorNameForeground,
		},
		Text: "Hello, World!",
	})
	md := widget.NewRichTextFromMarkdown(`
# Hello, World!

- This is a rich text widget.
- This is a rich text widget.

[https://hankmo.com](https://hankmo.com)
`)
	w.SetContent(container.New(layout.NewHBoxLayout(), richText, md))
	w.Resize(fyne.NewSize(200, 100))
	w.Show()
}

// Select is a widget that allows the user to select an option from a list of options.
func selectWidget(a fyne.App) {
	w := a.NewWindow("Select")
	selectWidget := widget.NewSelect([]string{"Option 1", "Option 2", "Option 3"}, func(s string) {
		println("Select value:", s)
	})
	w.SetContent(selectWidget)
	w.Resize(fyne.NewSize(200, 100))
	w.Show()
}

// SelectEntry is an input field which supports selecting from a fixed set of options.
// 可以输入，也可以选择
func selectEntry(a fyne.App) {
	w := a.NewWindow("SelectEntry")
	selectEntry := widget.NewSelectEntry([]string{"Option 1", "Option 2", "Option 3"})
	selectEntry.SetPlaceHolder("Select an option")
	selectEntry.OnChanged = func(s string) {
		println("SelectEntry value:", s)
	}
	w.SetContent(selectEntry)
	w.Resize(fyne.NewSize(200, 100))
	w.Show()
}

func separator(a fyne.App) {
	w := a.NewWindow("Separator")

	// 水平分离器示例
	button1 := widget.NewButton("Button 1", func() {
		println("Button 1 clicked")
	})
	button2 := widget.NewButton("Button 2", func() {
		println("Button 2 clicked")
	})
	button3 := widget.NewButton("Button 3", func() {
		println("Button 3 clicked")
	})
	button4 := widget.NewButton("Button 4", func() {
		println("Button 4 clicked")
	})
	button5 := widget.NewButton("Button 5", func() {
		println("Button 5 clicked")
	})

	// 垂直分离器示例
	label1 := widget.NewLabel("上侧内容")
	label2 := widget.NewLabel("下侧内容")

	// 使用不同的布局来展示分离器
	// 水平布局：按钮 - 分离器 - 按钮
	horizontalLayout := container.NewHBox(button1, widget.NewSeparator(), button2)

	// 垂直布局：标签 - 分离器 - 标签
	verticalLayout := container.NewVBox(label1, widget.NewSeparator(), label2)

	// 混合布局：按钮 - 分离器 - 按钮 - 分离器 - 按钮
	mixedLayout := container.NewVBox(
		button3,
		widget.NewSeparator(),
		button4,
		widget.NewSeparator(),
		button5,
	)

	// 将所有示例放在一个容器中
	content := container.NewVBox(
		widget.NewLabel("水平分离器:"),
		horizontalLayout,
		widget.NewLabel("垂直分离器:"),
		verticalLayout,
		widget.NewLabel("混合布局:"),
		mixedLayout,
	)

	w.SetContent(content)
	// w.Resize(fyne.NewSize(400, 300))
	w.Show()
}
