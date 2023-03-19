package shax

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha384Hex(data []byte) string {
	return hex.EncodeToString(Sha384(data))
}

func Sha384(data []byte) []byte {
	digest := sha512.New384()
	digest.Write(data)
	return digest.Sum(nil)
}
