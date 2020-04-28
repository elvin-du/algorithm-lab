package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 1, q.Size())
	q.Push(2)
	assert.Equal(t, 2, q.Peek())
	q.Push(3)
	assert.Equal(t, 3, q.Pop())
	assert.Equal(t, 2, q.Peek())
}
