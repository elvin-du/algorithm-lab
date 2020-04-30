package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	bt := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	root.SetRight(NewBinaryTreeNode(3))
	root.SetLeft(NewBinaryTreeNode(2))

	root.Left.SetLeft(NewBinaryTreeNode(4))
	root.Left.SetRight(NewBinaryTreeNode(5))

	root.Right.SetLeft(NewBinaryTreeNode(6))
	root.Right.SetRight(NewBinaryTreeNode(7))
	bt.Root = root

	assert.Equal(t, "1,2,3,4,5,6,7", bt.BFS())
	assert.Equal(t, "1245367", bt.PreOrder())
}
