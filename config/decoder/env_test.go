package decoder_test

import (
	"codeup.aliyun.com/qimao/leo/leo/config/decoder"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestENVIsSupported(t *testing.T) {
	envDec := decoder.ENV{}
	tests := []struct {
		Expected bool
		ext      string
	}{
		{
			Expected: true,
			ext:      "env",
		},
		{
			Expected: true,
			ext:      "Env",
		},
		{
			Expected: true,
			ext:      "ENv",
		},
		{
			Expected: true,
			ext:      "ENV",
		},
		{
			Expected: true,
			ext:      ".env",
		},
		{
			Expected: true,
			ext:      ".Env",
		},
		{
			Expected: true,
			ext:      ".ENv",
		},
		{
			Expected: true,
			ext:      ".ENV",
		},
		{
			Expected: false,
			ext:      "E.NV",
		},
		{
			Expected: false,
			ext:      "",
		},
		{
			Expected: false,
			ext:      "json",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.Expected, envDec.IsSupported(test.ext))
	}
}

var equalsEnv = `export OPTION_A='postgres://localhost:5432/database?sslmode=disable'
`
var exportedEnv = `export OPTION_A=2
export OPTION_B='\n'
`
var plainEnv = `OPTION_A=1
OPTION_B=2
OPTION_C= 3
OPTION_D =4
OPTION_E = 5
OPTION_F = 
OPTION_G=
OPTION_H=1 2`
var quotedEnv = `OPTION_A='1'
OPTION_B='2'
OPTION_C=''
OPTION_D='\n'
OPTION_E="1"
OPTION_F="2"
OPTION_G=""
OPTION_H="\n"
OPTION_I = "echo 'asd'"
OPTION_J='line 1
line 2'
OPTION_K='line one
this is \'quoted\'
one more line'
OPTION_L="line 1
line 2"
OPTION_M="line one
this is \"quoted\"
one more line"
`

func TestENVDecode(t *testing.T) {
	envDec := decoder.ENV{}
	tests := []string{equalsEnv, exportedEnv, plainEnv, quotedEnv}
	for _, test := range tests {
		m := map[string]any{}
		assert.NoError(t, envDec.Decode([]byte(test), m))
		m2, err := godotenv.UnmarshalBytes([]byte(test))
		assert.NoError(t, err)
		for key, val := range m {
			assert.Equal(t, val, m2[key])
		}
	}
}
