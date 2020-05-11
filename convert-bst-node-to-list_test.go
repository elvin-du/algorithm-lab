package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSTConvertToList(t *testing.T) {
	root := NewBinaryTreeNode(5)
	child2 := NewBinaryTreeNode(2)
	child3 := NewBinaryTreeNode(1)
	child4 := NewBinaryTreeNode(3)

	root.SetLeft(child2)
	child2.SetLeft(child3)
	child2.SetRight(child4)

	child5 := NewBinaryTreeNode(8)
	child6 := NewBinaryTreeNode(7)
	child7 := NewBinaryTreeNode(9)

	child5.SetLeft(child6)
	child5.SetRight(child7)

	root.SetRight(child5)

	newHead := BSTConvertToList(root)
	assert.Equal(t, 1, newHead.Value)
}
