package filex

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExist(t *testing.T) {
	f, err := os.CreateTemp("", "_Go_ErrIsExist")
	assert.Nil(t, err)
	defer os.Remove(f.Name())
	defer f.Close()

	dir := filepath.Dir(f.Name())
	isExist := IsExist(dir)
	assert.True(t, isExist)

	isExist = IsExist(dir + "tmp")
	assert.False(t, isExist)

	isExist = IsExist(f.Name())
	assert.True(t, isExist)

	isExist = IsExist(filepath.Join(dir, "_Go_ErrIsNotExist"))
	assert.False(t, isExist)

}
