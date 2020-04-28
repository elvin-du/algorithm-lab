package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyStack(t *testing.T) {
	mystack := NewMyStack()
	assert.Equal(t, true, mystack.Empty())
	mystack.Push(1)
	assert.Equal(t, false, mystack.Empty())
	assert.Equal(t, 1, mystack.Top())
	mystack.Push(2)
	mystack.Push(3)
	assert.Equal(t, 3, mystack.Top())
	assert.Equal(t, 3, mystack.Pop())
	assert.Equal(t, 2, mystack.Pop())
	assert.Equal(t, false, mystack.Empty())
}
