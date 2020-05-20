package code_interviews

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBalancedTree(t *testing.T) {
	root := common.NewBinaryTreeNode(1)
	child2 := common.NewBinaryTreeNode(2)
	child3 := common.NewBinaryTreeNode(3)

	root.SetLeft(child2)
	root.SetRight(child3)

	child4 := common.NewBinaryTreeNode(4)
	child5 := common.NewBinaryTreeNode(5)
	child2.SetRight(child4)
	child3.SetRight(child5)

	assert.Equal(t,true,IsBalancedTree(root))
}