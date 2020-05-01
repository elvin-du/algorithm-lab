package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNextNode(t *testing.T) {
	bt := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	child2 := NewBinaryTreeNode(2)
	child3 := NewBinaryTreeNode(3)

	bt.Root = root
	root.SetLeft(child2)
	root.SetRight(child3)

	child4 := NewBinaryTreeNode(4)
	child5 := NewBinaryTreeNode(5)
	child2.SetRight(child4)
	child3.SetRight(child5)

	child8 := NewBinaryTreeNode(8)
	child2.SetLeft(child8)

	assert.Equal(t, 1, GetNextNode(child4).Value)
	assert.Equal(t, 4, GetNextNode(child2).Value)
	assert.Equal(t, 4, GetNextNode(child2).Value)
	assert.Equal(t, 2, GetNextNode(child8).Value)
	assert.EqualValues(t, NilBinaryTreeNode, GetNextNode(child5))
}
