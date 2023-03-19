package hmacx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSha256(key []byte, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func HmacSha256Hex(key []byte, data []byte) string {
	return hex.EncodeToString(HmacSha256(key, data))
}
