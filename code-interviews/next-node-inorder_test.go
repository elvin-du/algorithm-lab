package code_interviews

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextNodeInorder(t *testing.T) {
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

	assert.Equal(t, 1, NextNodeInorder(child4).Value)
	assert.Equal(t, 4, NextNodeInorder(child2).Value)
	assert.Equal(t, 4, NextNodeInorder(child2).Value)
	assert.Equal(t, 2, NextNodeInorder(child8).Value)
	assert.EqualValues(t, common.NilBinaryTreeNode, NextNodeInorder(child5))
}