package shax

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha224Hex(data []byte) string {
	return hex.EncodeToString(Sha256(data))
}

func Sha224(data []byte) []byte {
	digest := sha256.New224()
	digest.Write(data)
	return digest.Sum(nil)
}
