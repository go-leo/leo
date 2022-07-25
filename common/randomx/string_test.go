package randomx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-leo/leo/common/randomx"
)

func TestNumericPermString(t *testing.T) {
	permString := randomx.NumericString(10)
	assert.Len(t, permString, 10)
	permString = randomx.NumericString(15)
	assert.Len(t, permString, 15)
	permString = randomx.NumericString(20)
	assert.Len(t, permString, 20)
	permString = randomx.NumericString(21)
	assert.Len(t, permString, 21)
	permString = randomx.NumericString(30)
	assert.Len(t, permString, 30)
	permString = randomx.NumericString(39)
	assert.Len(t, permString, 39)
}

func TestWordString(t *testing.T) {
	permString := randomx.WordString(10)
	assert.Len(t, permString, 10)
	permString = randomx.WordString(15)
	assert.Len(t, permString, 15)
	permString = randomx.WordString(20)
	assert.Len(t, permString, 20)
	permString = randomx.WordString(21)
	assert.Len(t, permString, 21)
	permString = randomx.WordString(30)
	assert.Len(t, permString, 30)
	permString = randomx.WordString(39)
	assert.Len(t, permString, 39)
	permString = randomx.WordString(16)
	t.Log(permString)
}
