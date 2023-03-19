package rsax

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"

	"github.com/go-leo/gox/cryptox/base64x"
)

func GenerateKeyHex(bits int) (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	privateKeyStr := hex.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey))
	publicKeyStr := hex.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	return privateKeyStr, publicKeyStr, nil
}

func GenerateKeyBase64(bits int) (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	privateKeyStr := base64x.StdEncode(x509.MarshalPKCS1PrivateKey(privateKey))
	publicKeyStr := base64x.StdEncode(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	return privateKeyStr, publicKeyStr, nil
}
