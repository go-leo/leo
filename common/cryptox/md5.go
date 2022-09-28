package cryptox

import (
	"crypto/md5"
	"encoding/hex"
)

// Deprecated: Do not use. use github.com/go-leo/cryptox instead.
func MD5(data []byte) []byte {
	hash := md5.New()
	hash.Write(data)
	return hash.Sum(nil)
}

// Deprecated: Do not use. use github.com/go-leo/cryptox instead.
func MD5HexString(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
