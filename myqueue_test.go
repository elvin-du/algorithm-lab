package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyQueue(t *testing.T) {
	queue := NewMyQueue()

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, 2, queue.Pop())
	assert.Equal(t, false, queue.Empty())
}
