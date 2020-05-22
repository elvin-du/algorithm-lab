package leetcode

import (
	"algorithm-lab/common"
	"testing"
)

func Test_inorder(t *testing.T) {
	root := common.NewBinaryTreeNode(1)
	root.SetRight(common.NewBinaryTreeNode(3))
	root.SetLeft(common.NewBinaryTreeNode(2))

	root.Left.SetLeft(common.NewBinaryTreeNode(4))
	root.Left.SetRight(common.NewBinaryTreeNode(5))

	root.Right.SetLeft(common.NewBinaryTreeNode(6))
	root.Right.SetRight(common.NewBinaryTreeNode(7))

	inorder(root)
	preorder(root)
}

func Test_postOrder(t *testing.T) {
	root := common.NewBinaryTreeNode(1)
	root.SetRight(common.NewBinaryTreeNode(3))
	root.SetLeft(common.NewBinaryTreeNode(2))

	root.Left.SetLeft(common.NewBinaryTreeNode(4))
	root.Left.SetRight(common.NewBinaryTreeNode(5))

	root.Right.SetLeft(common.NewBinaryTreeNode(6))
	root.Right.SetRight(common.NewBinaryTreeNode(7))

	postOrder(root)
}
