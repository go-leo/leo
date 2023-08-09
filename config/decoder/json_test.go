package decoder_test

import (
	"codeup.aliyun.com/qimao/leo/leo/config/decoder"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONIsSupported(t *testing.T) {
	jsonEnc := decoder.JSON{}
	tests := []struct {
		Expected bool
		ext      string
	}{
		{
			Expected: true,
			ext:      "json",
		},
		{
			Expected: true,
			ext:      "JSON",
		},
		{
			Expected: true,
			ext:      ".json",
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
		assert.Equal(t, test.Expected, jsonEnc.IsSupported(test.ext))
	}
}

var arrayJson = `
	{"Name": "Platypus", "Order": "Monotremata"}
`

var rawJson = `{"precomputed": true}`

func TestJsonDecode(t *testing.T) {
	jsonDec := decoder.JSON{}
	tests := []string{arrayJson, rawJson}
	for _, test := range tests {
		m := map[string]any{}
		assert.NoError(t, jsonDec.Decode([]byte(test), m))
		m2 := map[string]any{}
		err := json.Unmarshal([]byte(test), &m2)
		assert.NoError(t, err)
		for key, val := range m {
			assert.Equal(t, val, m2[key])
		}
	}
}
