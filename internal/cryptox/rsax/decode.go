package rsax

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"

	"github.com/go-leo/gox/cryptox/base64x"
)

func DecodeKeyHex(data []byte) (string, string, error) {
	block, _ := pem.Decode(data)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", "", err
	}
	privateKeyStr := hex.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey))
	publicKeyStr := hex.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	return privateKeyStr, publicKeyStr, nil
}

func DecodeKeyBase64(data []byte) (string, string, error) {
	block, _ := pem.Decode(data)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", "", err
	}
	privateKeyStr := base64x.StdEncode(x509.MarshalPKCS1PrivateKey(privateKey))
	publicKeyStr := base64x.StdEncode(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	return privateKeyStr, publicKeyStr, nil
}
