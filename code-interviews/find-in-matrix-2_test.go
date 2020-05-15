package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindInMatrix2(t *testing.T) {
	matrix := [][]int{
		[]int{3, 5, 7, 9},
		[]int{5, 7, 9, 20},
	}

	assert.Equal(t, true, FindInMatrix2(7, matrix))
	assert.Equal(t, false, FindInMatrix2(1, matrix))
	assert.Equal(t, true, FindInMatrix2(20, matrix))
	assert.Equal(t, true, FindInMatrix2(3, matrix))
}
