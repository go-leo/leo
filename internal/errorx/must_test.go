package errorx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	v := Must(testString())
	assert.NotEmpty(t, v)
	v = Must(testStringError())
	assert.Empty(t, v)
}

func testString() (string, error) {
	return "testString", nil
}

func testStringError() (string, error) {
	return "", errors.New("error")
}
