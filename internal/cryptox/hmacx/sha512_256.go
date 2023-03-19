package hmacx

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func HmacSha512_256(key []byte, data []byte) []byte {
	h := hmac.New(sha512.New512_256, key)
	h.Write(data)
	return h.Sum(nil)
}

func HmacSha512_256Hex(key []byte, data []byte) string {
	return hex.EncodeToString(HmacSha512_256(key, data))
}
