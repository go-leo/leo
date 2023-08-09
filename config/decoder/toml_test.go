package decoder_test

import (
	"codeup.aliyun.com/qimao/leo/leo/config/decoder"
	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTOMLIsSupported(t *testing.T) {
	tomlEnc := decoder.JSON{}
	tests := []struct {
		Expected bool
		ext      string
	}{
		{
			Expected: true,
			ext:      "toml",
		},
		{
			Expected: true,
			ext:      "TOML",
		},
		{
			Expected: true,
			ext:      ".toml",
		},
		{
			Expected: false,
			ext:      "",
		},
		{
			Expected: false,
			ext:      "yaml",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.Expected, tomlEnc.IsSupported(test.ext))
	}
}

var tomlBlob = `
[dishes.hamboogie]
name = "Hamboogie with fries"
price = 10.99

[[dishes.hamboogie.ingredients]]
name = "Bread Bun"

[[dishes.hamboogie.ingredients]]
name = "Lettuce"

[[dishes.hamboogie.ingredients]]
name = "Real Beef Patty"

[[dishes.hamboogie.ingredients]]
name = "Tomato"

[dishes.eggsalad]
name = "Egg Salad with rice"
price = 3.99

[[dishes.eggsalad.ingredients]]
name = "Egg"

[[dishes.eggsalad.ingredients]]
name = "Mayo"

[[dishes.eggsalad.ingredients]]
name = "Rice"
`

func TestTOMLDecode(t *testing.T) {
	tomlDec := decoder.TOML{}
	tests := []string{tomlBlob}
	for _, test := range tests {
		m := map[string]any{}
		assert.NoError(t, tomlDec.Decode([]byte(test), m))
		m2 := map[string]any{}
		err := toml.Unmarshal([]byte(test), &m2)
		assert.NoError(t, err)
		for key, val := range m {
			assert.Equal(t, val, m2[key])
		}
	}
}
