package md4x

import (
	"encoding/hex"

	"golang.org/x/crypto/md4"
)

func MD4(data []byte) []byte {
	hash := md4.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func MD4Hex(data []byte) string {
	return hex.EncodeToString(MD4(data))
}
