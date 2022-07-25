package filex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtName(t *testing.T) {
	name := ExtName("")
	assert.Equal(t, "", name)

	name = ExtName("config")
	assert.Equal(t, "", name)

	name = ExtName("config.yaml")
	assert.Equal(t, "yaml", name)

	name = ExtName(".conf")
	assert.Equal(t, "conf", name)
}
