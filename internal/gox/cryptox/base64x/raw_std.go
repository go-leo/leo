package base64x

import "encoding/base64"

func RawStdEncode(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
}

func RawStdDecode(s string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(s)
}
