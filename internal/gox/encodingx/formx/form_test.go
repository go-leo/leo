package formx

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

type Data struct {
	Name    string  `form:"name"`
	Age     int     `form:"age"`
	Height  float64 `form:"height"`
	IsChild bool    `form:"is_child"`
}

func TestMarshalDefault(t *testing.T) {
	data, err := Marshal(&Data{
		Name:    "jax",
		Age:     10,
		Height:  1.30,
		IsChild: true,
	})
	assert.NoError(t, err)
	t.Log(string(data))
}

type DataJson struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Height  float64 `json:"height"`
	IsChild bool    `json:"is_child"`
}

func TestMarshalDefaultJson(t *testing.T) {
	data, err := Marshal(&DataJson{
		Name:    "jax",
		Age:     10,
		Height:  1.30,
		IsChild: true,
	})
	assert.NoError(t, err)
	t.Log(string(data))
}

func TestMarshalDefaultMap(t *testing.T) {
	data, err := Marshal(map[string]any{
		"Name":    "jax",
		"Age":     10,
		"Height":  1.30,
		"IsChild": true,
	})
	assert.NoError(t, err)
	t.Log(url.QueryUnescape(string(data)))
}
