package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasSubtree(t *testing.T) {
	bt := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	child2 := NewBinaryTreeNode(2)
	child3 := NewBinaryTreeNode(3)
	child4 := NewBinaryTreeNode(4)
	child5 := NewBinaryTreeNode(5)

	realRoot := NewBinaryTreeNode(6)
	realRoot.SetRight(NewBinaryTreeNode(9))
	realRoot.SetLeft(root)

	bt.Root = realRoot
	root.SetLeft(child2)
	root.SetRight(child3)

	child2.SetLeft(child4)
	child3.SetRight(child5)

	bt2 := NewBinaryTree()
	root2 := NewBinaryTreeNode(1)
	root2_child2 := NewBinaryTreeNode(2)
	root2_child3 := NewBinaryTreeNode(3)
	root2_child4 := NewBinaryTreeNode(4)
	root2_child5 := NewBinaryTreeNode(5)

	bt2.Root = root2
	root2.SetLeft(root2_child2)
	root2.SetRight(root2_child3)

	root2_child2.SetLeft(root2_child4)
	root2_child3.SetRight(root2_child5)

	assert.Equal(t, true, HasSubtree(bt, bt))
}

func TestIsBTSame(t *testing.T) {
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

	bt2 := NewBinaryTree()
	root2 := NewBinaryTreeNode(1)
	root2_child2 := NewBinaryTreeNode(2)
	root2_child3 := NewBinaryTreeNode(3)
	root2_child4 := NewBinaryTreeNode(4)
	root2_child5 := NewBinaryTreeNode(5)

	bt2.Root = root2
	root2.SetLeft(root2_child2)
	root2.SetRight(root2_child3)

	root2_child2.SetLeft(root2_child4)
	root2_child3.SetRight(root2_child5)

	assert.Equal(t, true, IsBTSame(bt.Root, bt2.Root))

	root2_child5.SetRight(NewBinaryTreeNode(9))
	assert.Equal(t, false, IsBTSame(bt.Root, bt2.Root))

	assert.Equal(t, true, IsSubTree(bt2.Root, bt.Root))
}
