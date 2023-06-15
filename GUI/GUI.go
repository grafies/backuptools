package GUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	theme "gitee.com/grafies/goTheme"
)

func GUI() {
	a := app.New()                          //创建一个fyne应用
	a.Settings().SetTheme(&theme.MyTheme{}) //设置并引入字体,让程序可以显示中文
	w := a.NewWindow("backupTools")         //创建窗口名称
	if desk, ok := a.(desktop.App); ok {    //设置最小化托盘内容
		m := fyne.NewMenu("backupTools",
			fyne.NewMenuItem("主页面", func() {
				w.Show() //显示主页面
			})) //点击此按钮后,重新显示窗口
		desk.SetSystemTrayMenu(m) //使用捕获的系统托盘菜单
	}

	w.Resize(fyne.NewSize(480, 360)) //设置初始化窗口大小
	w.SetContent(container.NewVBox())
	w.Show() //显示主页面

	a.Run()
}

func Window(w *fyne.Window) {

}
