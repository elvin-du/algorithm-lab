package code_interviews

import (
	"algorithm-lab/common"
	"testing"
)

func TestFindPath(t *testing.T) {
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

	target = 7
	FindPath(bt.Root)
	t.Log(allpath)

	//t.Log(FindPath(bt, 9))
	//t.Log(FindPath(bt, 5))
}