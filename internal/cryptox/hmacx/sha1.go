package hmacx

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func HmacSha1(key []byte, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func HmacSha1Hex(key []byte, data []byte) string {
	return hex.EncodeToString(HmacSha1(key, data))
}
