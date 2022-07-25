package sortx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsc(t *testing.T) {
	ints := []int{3, 1, 6, 4, 2, 5}
	Asc(ints)
	assert.Equal(t, ints, []int{1, 2, 3, 4, 5, 6})

	float64s := []float64{3, 1, 6, 4, 2, 5}
	Asc(float64s)
	assert.Equal(t, float64s, []float64{1, 2, 3, 4, 5, 6})

	strings := []string{"3", "1", "6", "4", "2", "5"}
	Asc(strings)
	assert.Equal(t, strings, []string{"1", "2", "3", "4", "5", "6"})
}

func TestDesc(t *testing.T) {
	ints := []int{3, 1, 6, 4, 2, 5}
	Desc(ints)
	assert.Equal(t, ints, []int{6, 5, 4, 3, 2, 1})

	float64s := []float64{3, 1, 6, 4, 2, 5}
	Desc(float64s)
	assert.Equal(t, float64s, []float64{6, 5, 4, 3, 2, 1})

	strings := []string{"3", "1", "6", "4", "2", "5"}
	Desc(strings)
	assert.Equal(t, strings, []string{"6", "5", "4", "3", "2", "1"})
}
