package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int{6, 8, 1, 5, 4, 3}
	QuickSort(data)
	assert.Equal(t, []int{1, 3, 4, 5, 6, 8}, data)
}
