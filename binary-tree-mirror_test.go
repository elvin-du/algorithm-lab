package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeMirror(t *testing.T) {
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

	assert.Equal(t, "1,2,3,8,4,5", bt.BFS())
	BinaryTreeMirror(bt)
	assert.Equal(t, "1,3,2,5,4,8", bt.BFS())
}
