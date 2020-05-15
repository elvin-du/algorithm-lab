package algorithm_lab

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeDepth(t *testing.T) {
	root := common.NewBinaryTreeNode(1)
	child2 := common.NewBinaryTreeNode(2)
	child3 := common.NewBinaryTreeNode(2)
	child4 := common.NewBinaryTreeNode(2)
	child5 := common.NewBinaryTreeNode(2)
	child6 := common.NewBinaryTreeNode(2)

	root.Left = child2
	root.Right = child3
	child2.Left = child4
	child3.Right = child5
	child5.Left = child6

	assert.Equal(t, 4, TreeDepth(root))
}
