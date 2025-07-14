# 打包

图标文件默认为Icon.png, 否则需要`-icon`参数传递图标文件，fyne会自动转换。

## 桌面打包

```bash
fyne package -os darwin -icon myapp.png
fyne package -os windows -icon myapp.png
fyne package -os linux -icon myapp.png
```

将创建:

- myapp.tar.gz，其中包含从 usr/local/ 开始的文件夹结构，Linux 用户可以将其扩展到其系统的根目录。
- myapp.exe（在第二次构建之后，这是 Windows 包所必需的）将嵌入图标和应用程序元数据。

## 移动打包

```bash
fyne package -os android -app-id com.example.myapp -icon mobileIcon.png
fyne package -os ios -app-id com.example.myapp -icon mobileIcon.png
```

然后在Android和模拟器中安装:

```bash
adb install myapp.apk
```

要在 iOS 设备上安装，请打开 Xcode，然后在“窗口”菜单中选择“设备和模拟器”菜单项。然后找到您的手机并将 myapp.app 图标拖到您的应用列表中。

如果要在模拟器上安装，请确保使用 iossimulator 与 ios 打包应用程序:

```bash
fyne package -os iossimulator -app-id com.example.myapp -icon mobileIcon.png
xcrun simctl install booted myapp.app
```

# 在浏览器中运行

Fyne 应用程序也可以使用标准 Web 浏览器在 Web 上运行！使用 Fyne 创建的 Web 应用程序将捆绑一个 WebAssembly 二进制文件，该二进制文件将在大多数现代浏览器中运行。

为了准备您的应用程序以在 Web 上使用，我们再次使用 “fyne” cli 应用程序，它有一个用于快速测试的 “serve” 命令：

```bash
go install fyne.io/tools/cmd/fyne@latest
fyne serve -icon icon.png
```

然后浏览器访问: `https://localhost:8080`

`fyne serve` 命令非常适合本地测试，但就像您需要的其他平台一样 以便能够分发您的应用程序。要准备要上传的文件，只需使用 `fyne package` 命令，就像常规的 Packaging 一样。

```bash
fyne package -os web
```

# Metadata元数据

Metadata元数据文件有助于避免记住每个 package 和 release 命令的特定构建参数。

## 基本配置

该文件应在运行 fyne 命令的目录中命名为 FyneApp.toml（这通常是主包）。该文件的内容如下：

```toml
Website = "https://example.com"

[Details]
Icon = "Icon.png"
Name = "My App"
ID = "com.example.app"
Version = "1.0.0"
Build = 1
```

文件的顶部是元数据，如果您将应用程序上传到 Apps 列表页面，将使用该元数据，因此它是可选的。

Details （详细信息 ） 表包含有关您的应用程序的数据，这些数据由其他应用商店和作系统在发布过程中使用。

如果找到此文件，fyne 工具将使用此文件，如果存在元数据，则不需要许多强制性命令参数。您仍然可以使用命令行参数覆盖这些值。

## Linux & BSD 配置

`LinuxAndBSD` 的可选表包含 `Fyne` 应用程序的 “desktop entry” 配置文件的其他参数。所有参数都是可选的，但如果存在， 则 fyne 工具将使用它们（除了 Details 表中的参数之外）。

示例:

```toml
[LinuxAndBSD]
GenericName = "Web Browser"
Categories = ["Network"]
Comment = "View sites on the Internet"
Keywords = ["browser", "web"]
ExecParams = "-x 42"
```
