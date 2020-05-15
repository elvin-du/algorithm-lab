package algorithm_lab

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeMirror(t *testing.T) {
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

	assert.Equal(t, "1,2,3,8,4,5", bt.BFS())
	BinaryTreeMirror(bt)
	assert.Equal(t, "1,3,2,5,4,8", bt.BFS())
}
