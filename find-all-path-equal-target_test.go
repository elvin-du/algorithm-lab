package algorithm_lab

import (
	"algorithm-lab/common"
	"testing"
)

func TestFindAllPath(t *testing.T) {
	bt := common.NewBinaryTree()
	root := common.NewBinaryTreeNode(1)
	child2 := common.NewBinaryTreeNode(2)
	child3 := common.NewBinaryTreeNode(3)
	child4 := common.NewBinaryTreeNode(4)
	child5 := common.NewBinaryTreeNode(5)

	bt.Root = root
	root.SetLeft(child2)
	root.SetRight(child3)

	child2.SetLeft(child4)
	child3.SetRight(child5)

	t.Log(FindAllPath(bt, 7))
	t.Log(FindAllPath(bt, 9))
	t.Log(FindAllPath(bt, 5))
}
