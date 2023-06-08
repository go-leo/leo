package base64x

import "encoding/base64"

func URLEncode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func URLDecode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}
