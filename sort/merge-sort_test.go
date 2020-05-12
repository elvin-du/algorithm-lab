package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := []int{8, 6, 1, 10, 22, 33, 9, 21, 5}
	assert.Equal(t, []int{1, 5, 6, 8, 9, 10, 21, 22, 33}, MergeSort(data))
}
