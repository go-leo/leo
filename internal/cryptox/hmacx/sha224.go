package hmacx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSha224(key []byte, data []byte) []byte {
	h := hmac.New(sha256.New224, key)
	h.Write(data)
	return h.Sum(nil)
}

func HmacSha224Hex(key []byte, data []byte) string {
	return hex.EncodeToString(HmacSha224(key, data))
}
