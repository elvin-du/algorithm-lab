package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExistInMatrix(t *testing.T) {
	matrix := [][]int{
		[]int{3, 5, 7, 9},
		[]int{5, 7, 9, 20},
	}

	assert.Equal(t, true, ExistInMatrix(7, matrix))
	assert.Equal(t, false, ExistInMatrix(1, matrix))
	assert.Equal(t, true, ExistInMatrix(20, matrix))
	assert.Equal(t, true, ExistInMatrix(3, matrix))
}
