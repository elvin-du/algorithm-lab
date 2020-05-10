package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReorderArray(t *testing.T) {
	data := []int{4, 5, 8, 7, 9, 1, 3, 6}
	ReorderArray(data)
	assert.Equal(t, []int{3, 5, 1, 7, 9, 8, 4, 6}, data)

	data = []int{4, 5, 8, 7, 9, 1, 3, 6, 33, 44, 100, 99, 123, 98, 78}
	ReorderArray(data)
	assert.Equal(t, []int{123, 5, 99, 7, 9, 1, 3, 33, 6, 44, 100, 8, 4, 98, 78}, data)
}
