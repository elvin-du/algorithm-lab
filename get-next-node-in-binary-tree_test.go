package algorithm_lab

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNextNode(t *testing.T) {
	bt := common.NewBinaryTree()
	root := common.NewBinaryTreeNode(1)
	child2 := common.NewBinaryTreeNode(2)
	child3 := common.NewBinaryTreeNode(3)

	bt.Root = root
	root.SetLeft(child2)
	root.SetRight(child3)

	child4 := common.NewBinaryTreeNode(4)
	child5 := common.NewBinaryTreeNode(5)
	child2.SetRight(child4)
	child3.SetRight(child5)

	child8 := common.NewBinaryTreeNode(8)
	child2.SetLeft(child8)

	assert.Equal(t, 1, GetNextNode(child4).Value)
	assert.Equal(t, 4, GetNextNode(child2).Value)
	assert.Equal(t, 4, GetNextNode(child2).Value)
	assert.Equal(t, 2, GetNextNode(child8).Value)
	assert.EqualValues(t, common.NilBinaryTreeNode, GetNextNode(child5))
}
