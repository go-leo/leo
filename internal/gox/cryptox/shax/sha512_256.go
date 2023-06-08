package shax

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512_256Hex(data []byte) string {
	return hex.EncodeToString(Sha512_256(data))
}

func Sha512_256(data []byte) []byte {
	digest := sha512.New512_256()
	digest.Write(data)
	return digest.Sum(nil)
}
