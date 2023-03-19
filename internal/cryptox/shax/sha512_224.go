package shax

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512_224Hex(data []byte) string {
	return hex.EncodeToString(Sha512_224(data))
}

func Sha512_224(data []byte) []byte {
	digest := sha512.New512_224()
	digest.Write(data)
	return digest.Sum(nil)
}
