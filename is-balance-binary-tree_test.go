package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBalancedBinaryTree(t *testing.T) {
	root := NewBinaryTreeNode(1)
	child2 := NewBinaryTreeNode(2)
	child3 := NewBinaryTreeNode(3)

	root.SetLeft(child2)
	root.SetRight(child3)

	child4 := NewBinaryTreeNode(4)
	child5 := NewBinaryTreeNode(5)
	child2.SetRight(child4)
	child3.SetRight(child5)

	assert.Equal(t,true,IsBalancedBinaryTree(root))
}