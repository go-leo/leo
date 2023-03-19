package hmacx

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func HmacSha384(key []byte, data []byte) []byte {
	h := hmac.New(sha512.New384, key)
	h.Write(data)
	return h.Sum(nil)
}

func HmacSha384Hex(key []byte, data []byte) string {
	return hex.EncodeToString(HmacSha384(key, data))
}
