package rsax

import (
	"os"
)

func LoadKeyHex(filename string) (string, string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", "", err
	}
	return DecodeKeyHex(data)
}

func LoadKeyBase64(filename string) (string, string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", "", err
	}
	return DecodeKeyBase64(data)
}
