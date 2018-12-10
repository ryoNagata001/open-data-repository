package util

import "fmt"

// byteコードをhexに変換する
func ByteToHex(input []byte) string {
	var hexValue string
	for _, v := range input {
		hexValue += fmt.Sprintf("%02x", v)
	}
	return hexValue
}