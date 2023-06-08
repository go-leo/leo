package hmacx

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func HmacMD5(key []byte, data []byte) []byte {
	h := hmac.New(md5.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func HmacMD5Hex(key []byte, data []byte) string {
	return hex.EncodeToString(HmacMD5(key, data))
}
