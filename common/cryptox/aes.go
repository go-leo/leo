package cryptox

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// pad 填充到BlockSize整数倍长度，如果正好就是对的长度，再多填充一个BlockSize长度
func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// unpad 去除填充的字节
func unpad(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return []byte{0}, nil
	}
	unpadding := int(src[length-1])
	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}
	return src[:(length - unpadding)], nil
}

// AESEncrypt 加密
// Deprecated: Do not use. use github.com/go-leo/cryptox instead.
func AESEncrypt(text string, Key string, IV string) (string, error) {
	if "" == text {
		return "", nil
	}
	block, err := aes.NewCipher([]byte(Key))
	if err != nil {
		return "", err
	}
	msg := pad([]byte(text))
	ciphertext := make([]byte, len(msg))
	mode := cipher.NewCBCEncrypter(block, []byte(IV))
	mode.CryptBlocks(ciphertext, msg)
	finalMsg := base64.StdEncoding.EncodeToString(ciphertext)
	return finalMsg, nil
}

// AESDecrypt 解密
// Deprecated: Do not use. use github.com/go-leo/cryptox instead.
func AESDecrypt(text string, key string, iv string) (string, error) {
	if "" == text {
		return "", nil
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decodedMsg, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blocksize must be multipe of decoded message length")
	}
	msg := decodedMsg
	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(msg, msg)
	unpadMsg, err := unpad(msg)
	if err != nil {
		return "", err
	}
	return string(unpadMsg), nil
}
