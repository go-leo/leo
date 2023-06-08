package base64x

import "encoding/base64"

func RawURLEncode(src []byte) string {
	return base64.RawURLEncoding.EncodeToString(src)
}

func RawURLDecode(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
