package mathx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, uint(1), Min(uint(1), uint(2)))
	assert.Equal(t, uint8(1), Min(uint8(1), uint8(2)))
	assert.Equal(t, uint16(1), Min(uint16(1), uint16(2)))
	assert.Equal(t, uint32(1), Min(uint32(1), uint32(2)))
	assert.Equal(t, uint64(1), Min(uint64(1), uint64(2)))
	assert.Equal(t, int(1), Min(int(1), int(2)))
	assert.Equal(t, int8(1), Min(int8(1), int8(2)))
	assert.Equal(t, int16(1), Min(int16(1), int16(2)))
	assert.Equal(t, int32(1), Min(int32(1), int32(2)))
	assert.Equal(t, int64(1), Min(int64(1), int64(2)))
	assert.Equal(t, float32(1), Min(float32(1), float32(2)))
	assert.Equal(t, float64(1), Min(float64(1), float64(2)))
}
