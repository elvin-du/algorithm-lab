package algorithm_lab

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSTConvertToList(t *testing.T) {
	root := common.NewBinaryTreeNode(5)
	child2 := common.NewBinaryTreeNode(2)
	child3 := common.NewBinaryTreeNode(1)
	child4 := common.NewBinaryTreeNode(3)

	root.SetLeft(child2)
	child2.SetLeft(child3)
	child2.SetRight(child4)

	child5 := common.NewBinaryTreeNode(8)
	child6 := common.NewBinaryTreeNode(7)
	child7 := common.NewBinaryTreeNode(9)

	child5.SetLeft(child6)
	child5.SetRight(child7)

	root.SetRight(child5)

	newHead := BSTConvertToList(root)
	assert.Equal(t, 1, newHead.Value)
}
