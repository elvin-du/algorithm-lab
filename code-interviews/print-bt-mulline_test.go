package code_interviews

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintMulLine(t *testing.T) {
	tree := common.NewBinaryTree()
	root := common.NewBinaryTreeNode(1)
	tree.Root = root

	child1 := common.NewBinaryTreeNode(2)
	child2 := common.NewBinaryTreeNode(2)
	child3 := common.NewBinaryTreeNode(4)
	child4 := common.NewBinaryTreeNode(3)
	child5 := common.NewBinaryTreeNode(3)
	child6 := common.NewBinaryTreeNode(4)

	root.SetLeft(child1)
	root.SetRight(child2)

	child1.SetLeft(child3)
	child1.SetRight(child4)

	child2.SetLeft(child5)
	child2.SetRight(child6)

	assert.Equal(t,[]string{"1", "22", "4334"},PrintMulLine(tree.Root))
}