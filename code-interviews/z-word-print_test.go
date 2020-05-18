package code_interviews

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZWordPrint(t *testing.T) {
	tree := common.NewBinaryTree()
	root := common.NewBinaryTreeNode(1)
	child2 := common.NewBinaryTreeNode(2)
	child3 := common.NewBinaryTreeNode(3)
	child4 := common.NewBinaryTreeNode(4)
	child5 := common.NewBinaryTreeNode(5)
	child6 := common.NewBinaryTreeNode(6)
	child7 := common.NewBinaryTreeNode(7)
	child8 := common.NewBinaryTreeNode(8)
	child9 := common.NewBinaryTreeNode(9)
	child10 := common.NewBinaryTreeNode(10)
	child11 := common.NewBinaryTreeNode(11)
	child12 := common.NewBinaryTreeNode(12)

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

	assert.Equal(t, "1 3 2 4 5 6 7 12 11 10 9 8 ", ZWordPrint(tree.Root))
	child7.SetLeft(common.NewBinaryTreeNode(14))

	assert.Equal(t, "1 3 2 4 5 6 7 14 12 11 10 9 8 ", ZWordPrint(tree.Root))
}
