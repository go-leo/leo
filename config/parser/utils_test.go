package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardized_int(t *testing.T) {
	v := standardized(1)
	assert.Equal(t, 1, v)
}

func TestStandardized_float(t *testing.T) {
	v := standardized(1.2)
	assert.Equal(t, 1.2, v)
}

func TestStandardized_string(t *testing.T) {
	v := standardized("this is string")
	assert.Equal(t, "this is string", v)
}

func TestStandardized_bool(t *testing.T) {
	v := standardized(true)
	assert.Equal(t, true, v)
}

func TestStandardized_intSlice(t *testing.T) {
	v := standardized([]int{1, 2})
	assert.Equal(t, []any{1, 2}, v)
}

func TestStandardized_stringSlice(t *testing.T) {
	v := standardized([]string{"1", "2"})
	assert.Equal(t, []any{"1", "2"}, v)
}

func TestStandardized_boolSlice(t *testing.T) {
	v := standardized([]bool{true, false})
	assert.Equal(t, []any{true, false}, v)
}

func TestStandardized_intMap(t *testing.T) {
	v := standardized(map[string]int{
		"k1": 1,
		"k2": 2,
	})
	assert.Equal(t, map[string]any{
		"k1": 1,
		"k2": 2,
	}, v)
}

func TestStandardized_floatMap(t *testing.T) {
	v := standardized(map[string]float64{
		"k1": 1.1,
		"k2": 2.2,
	})
	assert.Equal(t, map[string]any{
		"k1": 1.1,
		"k2": 2.2,
	}, v)
}

func TestStandardized_stringMap(t *testing.T) {
	v := standardized(map[string]string{
		"k1": "1.1",
		"k2": "2.2",
	})
	assert.Equal(t, map[string]any{
		"k1": "1.1",
		"k2": "2.2",
	}, v)
}

func TestStandardized_boolMap(t *testing.T) {
	v := standardized(map[string]bool{
		"k1": true,
		"k2": false,
	})
	assert.Equal(t, map[string]any{
		"k1": true,
		"k2": false,
	}, v)
}

func TestStandardized_MapMap(t *testing.T) {
	v := standardized(map[string]any{
		"k1": map[string]any{
			"k1k1": 1,
			"k1k2": 2,
		},
		"k2": []string{"1", "2"},
	})
	assert.Equal(t, map[string]any{
		"k1": map[string]any{
			"k1k1": 1,
			"k1k2": 2,
		},
		"k2": []any{"1", "2"},
	}, v)
}
