package GUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gitee.com/grafies/goTypeface"
	"log"
)

func GUI() {
	a := app.New()                                  //创建一个fyne应用
	a.Settings().SetTheme(&goTypeface.GoTypeface{}) //设置并引入字体,让程序可以显示中文
	w := a.NewWindow("backupTools")                 //创建窗口名称

	MainShow(w) //调用窗口函数

	if desk, ok := a.(desktop.App); ok { //设置最小化托盘内容
		m := fyne.NewMenu("backupTools",
			fyne.NewMenuItem("主页面", func() {
				w.Show() //显示主页面
			}),
			fyne.NewMenuItem("设置", func() {

			})) //点击此按钮后,重新显示窗口
		desk.SetSystemTrayMenu(m) //使用捕获的系统托盘菜单
	}

	w.Resize(fyne.NewSize(800, 600)) //设置初始化窗口大小
	w.CenterOnScreen()               //居中显示

	//w.SetContent(container.NewVBox())
	w.ShowAndRun() //显示主页面

}

func MainShow(w fyne.Window) {
	title := widget.NewLabel("备份工具")
	compressedDirectory := widget.NewLabel("压缩目录：")
	sourceAddress1 := widget.NewEntry() //输入的文本框

	dia1 := widget.NewButton("浏览", func() { //回调函数：打开选择文件对话框
		fd := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if closer == nil {
				log.Println("取消")

			}
			sourceAddress1.SetText(closer.URI().Path())
		}, w)

		//fd.SetFilter(storage.NewExtensionFileFilter([]string{".txt"})) //打开的文件格式类型
		fd.Show() //显示控件
	})

	text := widget.NewMultiLineEntry() //多行输入组件
	//text.Disable()                     //禁用输入框，不能更改数据

	storageDirectory := widget.NewLabel("存放目录：")
	sourceAddress2 := widget.NewEntry() //输入的文本框

	dia2 := widget.NewButton("浏览", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if uri != nil {
				log.Println("取消")

			}

			sourceAddress2.SetText(uri.Path()) //设置输入框内容
		}, w)
	})

	dia3 := widget.NewButton("开始压缩", func() {
		go func() {
			//获取内容
			if compressedDirectory.Text != "" && storageDirectory.Text != "" {
				text.SetText("")
				text.Refresh()
			}
		}()
	})

	head := container.NewCenter(title)

	//第一行和第二行
	v1 := container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), compressedDirectory, dia1, sourceAddress1)
	v2 := container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), storageDirectory, dia2, sourceAddress2)

	v3 := container.NewHBox(dia3)       //显示第三行按钮
	v3Center := container.NewCenter(v3) //第三行

	labelLast := widget.NewLabel("backupTools")

	position := container.NewVBox(head, v1, v2, v3Center, text, labelLast) //控制显示位置顺序

	w.SetContent(position)
}
