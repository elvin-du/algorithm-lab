package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapSort(t *testing.T) {
	data := []int{5, 7, 1, 2, 22, 33, 6, 2}
	HeapSort(data)

	assert.Equal(t, []int{1, 2, 2, 5, 6, 7, 22, 33}, data)
}

func TestInsertMaxHeap(t *testing.T) {
	data := []int{5, 7, 1, 2, 22, 33, 6, 2}
	HeapSort(data)

	assert.Equal(t, []int{1, 2, 2, 5, 6, 7, 22, 33}, data)

	data = InsertMaxHeap(data, 8)
	HeapSort(data)
	assert.Equal(t, []int{1, 2, 2, 5, 6, 7, 8, 22, 33}, data)

	data = InsertMaxHeap(data, 88)
	assert.Equal(t, 88, data[0])
}

func TestMinHeapSort(t *testing.T) {
	data := []int{5, 7, 1, 2, 22, 33, 6, 2}
	MinHeapSort(data)

	assert.Equal(t, []int{33, 22, 7, 6, 5, 2, 2, 1}, data)
}

func TestMaxHeapPoll(t *testing.T) {
	data := []int{5, 7, 1, 2, 22, 33, 6, 2}
	BuildMaxHeap(data, len(data))
	assert.Equal(t, 33, data[0])
	assert.Equal(t, 33, MaxHeapPoll(&data))
	assert.Equal(t, 7, len(data))
}
