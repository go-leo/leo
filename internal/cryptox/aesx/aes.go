package aesx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"github.com/go-leo/gox/cryptox/base64x"
)

// Encrypt 加密
func Encrypt(text string, Key string, IV string) (string, error) {
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
	finalMsg := base64x.StdEncode(ciphertext)
	return finalMsg, nil
}

// Decrypt 解密
func Decrypt(text string, key string, iv string) (string, error) {
	if "" == text {
		return "", nil
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decodedMsg, err := base64x.StdDecode(text)
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

// pad 填充到BlockSize整数倍长度，如果正好就是对的长度，再多填充一个BlockSize长度
func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
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
