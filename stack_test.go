package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack(100)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	assert.Equal(t, 3, stack.Size())
	assert.Equal(t, 3, stack.Top())
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Top())
}
