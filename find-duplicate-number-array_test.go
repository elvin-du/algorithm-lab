package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicateNumberInArray(t *testing.T) {
	arr := []int{2, 3, 1, 0, 2, 5, 3}
	assert.Equal(t, 2, FindDuplicateNumberInArray(arr))

	arr2 := []int{3, 1, 0, 2, 5, 3}
	assert.Equal(t, 3, FindDuplicateNumberInArray(arr2))
}
