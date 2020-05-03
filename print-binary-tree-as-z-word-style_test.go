package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintBinaryTreeAsZWordStyle(t *testing.T) {
	tree := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	child2 := NewBinaryTreeNode(2)
	child3 := NewBinaryTreeNode(3)
	child4 := NewBinaryTreeNode(4)
	child5 := NewBinaryTreeNode(5)
	child6 := NewBinaryTreeNode(6)
	child7 := NewBinaryTreeNode(7)
	child8 := NewBinaryTreeNode(8)
	child9 := NewBinaryTreeNode(9)
	child10 := NewBinaryTreeNode(10)
	child11 := NewBinaryTreeNode(11)
	child12 := NewBinaryTreeNode(12)

	tree.Root = root
	root.SetLeft(child2)
	root.SetRight(child3)

	child2.SetLeft(child4)
	child2.SetRight(child5)

	child3.SetLeft(child6)
	child3.SetRight(child7)

	child4.SetLeft(child8)
	child4.SetRight(child9)

	child5.SetLeft(child10)
	child5.SetRight(child11)

	child6.SetLeft(child12)

	assert.Equal(t, "1 3 2 4 5 6 7 12 11 10 9 8 ", PrintBinaryTreeAsZWordStyle(tree))
	child7.SetLeft(NewBinaryTreeNode(14))

	assert.Equal(t, "1 3 2 4 5 6 7 14 12 11 10 9 8 ", PrintBinaryTreeAsZWordStyle(tree))
}
