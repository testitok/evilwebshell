package main

import (
	_ "embed"
	"evilwebshell/core"
	"evilwebshell/themes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/dialog"
	"os"
)

const preferenceCurrentTutorial = "currentTutorial"
const Version = "1.3.0"

var topWindow fyne.Window

// 设置中文
func init() {

}

func main() {
	defer os.Unsetenv("FYNE_FONT")
	//退出时
	defer os.RemoveAll(core.TempDir)
	a := app.NewWithID("io.fyne.demo")
	a.SetIcon(themes.Resource2Png)
	//a.SetIcon(theme.FyneLogo())
	//a := app.New() //新建一个应用
	w := a.NewWindow("evilwebshell") //新建一个窗口
	settingsItem := fyne.NewMenuItem("Settings", func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	})
	version := fyne.NewMenuItem("VERSION", func() {
		v1 := dialog.NewInformation("Version", Version, w)
		v1.Show()
	})
	author := fyne.NewMenuItem("Author", func() {
		//author := a.NewWindow("piiperxyz")
		v2 := dialog.NewInformation("Author", "evilwebshell\nhttps://github.com/evilwebshell", w)
		v2.Show()
	})
	info := fyne.NewMenuItem("Info", func() {
		//author := a.NewWindow("piiperxyz")
		v3 := dialog.NewInformation("Info", "from\nhttps://github.com/piiperxyz/AniYa", w)
		v3.Show()
	})
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("FILE", settingsItem),
		fyne.NewMenu("ABOUT", version, author, info),
	)
	tmp := themes.BypassAV(w)

	w.SetContent(tmp)
	w.SetMainMenu(mainMenu)
	w.SetMaster()
	w.Resize(fyne.NewSize(800, 700))
	w.ShowAndRun() //显示窗口并运行，后续的窗口只能用show
	w.Show()       //显示窗口并运行，后续的窗口只能用show
}
