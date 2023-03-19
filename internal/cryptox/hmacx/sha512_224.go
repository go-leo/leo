package hmacx

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func HmacSha512_224(key []byte, data []byte) []byte {
	h := hmac.New(sha512.New512_224, key)
	h.Write(data)
	return h.Sum(nil)
}

func HmacSha512_224Hex(key []byte, data []byte) string {
	return hex.EncodeToString(HmacSha512_224(key, data))
}
