package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSymmetricalBinaryTree(t *testing.T) {
	tree := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	tree.Root = root

	child1 := NewBinaryTreeNode(2)
	child2 := NewBinaryTreeNode(2)
	child3 := NewBinaryTreeNode(3)
	child4 := NewBinaryTreeNode(3)
	child5 := NewBinaryTreeNode(4)
	child6 := NewBinaryTreeNode(4)

	root.SetLeft(child1)
	root.SetRight(child2)

	child1.SetLeft(child3)
	child1.SetRight(child4)

	child2.SetLeft(child5)
	child2.SetRight(child6)

	assert.Equal(t, true, IsSymmetricalBinaryTree(tree))
	child6.SetRight(NewBinaryTreeNode(8))
	assert.Equal(t, false, IsSymmetricalBinaryTree(tree))
}
