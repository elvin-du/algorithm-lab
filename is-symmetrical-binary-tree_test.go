package algorithm_lab

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSymmetricalBinaryTree(t *testing.T) {
	tree := common.NewBinaryTree()
	root := common.NewBinaryTreeNode(1)
	tree.Root = root

	child1 := common.NewBinaryTreeNode(2)
	child2 := common.NewBinaryTreeNode(2)
	child3 := common.NewBinaryTreeNode(3)
	child4 := common.NewBinaryTreeNode(3)
	child5 := common.NewBinaryTreeNode(4)
	child6 := common.NewBinaryTreeNode(4)

	root.SetLeft(child1)
	root.SetRight(child2)

	child1.SetLeft(child3)
	child1.SetRight(child4)

	child2.SetLeft(child5)
	child2.SetRight(child6)

	assert.Equal(t, true, IsSymmetricalBinaryTree(tree))
	child6.SetRight(common.NewBinaryTreeNode(8))
	assert.Equal(t, false, IsSymmetricalBinaryTree(tree))
}
