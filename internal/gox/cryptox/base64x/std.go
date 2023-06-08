package base64x

import "encoding/base64"

func StdEncode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func StdDecode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
