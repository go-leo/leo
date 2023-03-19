package rsax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignHex(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyHex(1024)
	assert.NoError(t, err)

	t.Log(privateKey)
	t.Log(publicKey)

	rawData := []byte("he is hello kitty")
	signHex, err := SignWithSha256Hex(rawData, privateKey)
	assert.NoError(t, err)

	err = VerifySignWithSha256Hex(rawData, signHex, publicKey)
	assert.NoError(t, err)
}

func TestSignBase64(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyBase64(1024)
	assert.NoError(t, err)

	t.Log(privateKey)
	t.Log(publicKey)

	rawData := []byte("he is hello kitty")
	signHex, err := SignWithSha256Base64(rawData, privateKey)
	assert.NoError(t, err)

	err = VerifySignWithSha256Base64(rawData, signHex, publicKey)
	assert.NoError(t, err)
}

func TestCrypt(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyHex(1024)
	assert.NoError(t, err)

	rawData := []byte("he is hello kitty")
	encryptDate, err := EncryptToHex(rawData, publicKey)
	assert.NoError(t, err)

	data, err := DecryptByHex(encryptDate, privateKey)
	assert.NoError(t, err)

	assert.Equal(t, rawData, data)
}

func TestLoad(t *testing.T) {
	priv, pub, err := LoadKeyBase64("/tmp/priv.pem")
	assert.NoError(t, err)
	t.Log(priv)
	t.Log(pub)
}
