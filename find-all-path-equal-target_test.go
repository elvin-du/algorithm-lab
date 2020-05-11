package algorithm_lab

import (
	"testing"
)

func TestFindAllPath(t *testing.T) {
	bt := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	child2 := NewBinaryTreeNode(2)
	child3 := NewBinaryTreeNode(3)
	child4 := NewBinaryTreeNode(4)
	child5 := NewBinaryTreeNode(5)

	bt.Root = root
	root.SetLeft(child2)
	root.SetRight(child3)

	child2.SetLeft(child4)
	child3.SetRight(child5)

	t.Log(FindAllPath(bt, 7))
	t.Log(FindAllPath(bt, 9))
	t.Log(FindAllPath(bt, 5))
}
