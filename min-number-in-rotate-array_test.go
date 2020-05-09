package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinNumberInRotateArray(t *testing.T) {
	data := []int{3, 4, 5, 1, 2}
	assert.Equal(t, 1, MinNumberInRotateArray(data))

	data = []int{7, 8, 9, 3, 4, 5, 6}
	assert.Equal(t, 3, MinNumberInRotateArray(data))
}
