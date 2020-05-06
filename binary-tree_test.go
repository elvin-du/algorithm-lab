package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	bt := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	root.SetRight(NewBinaryTreeNode(3))
	root.SetLeft(NewBinaryTreeNode(2))

	root.Left.SetLeft(NewBinaryTreeNode(4))
	root.Left.SetRight(NewBinaryTreeNode(5))

	root.Right.SetLeft(NewBinaryTreeNode(6))
	root.Right.SetRight(NewBinaryTreeNode(7))
	bt.Root = root

	assert.Equal(t, "1,2,3,4,5,6,7", bt.BFS())
	assert.Equal(t, "1245367", bt.PreOrder())
}

func TestBinaryTree_BFS2(t *testing.T) {
	bt := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	root.SetLeft(NewBinaryTreeNode(2))
	root.SetRight(NewBinaryTreeNode(3))

	root.Left.SetLeft(NewBinaryTreeNode(4))
	root.Left.SetRight(NewBinaryTreeNode(5))

	root.Right.SetLeft(NewBinaryTreeNode(6))
	root.Right.SetRight(NewBinaryTreeNode(7))
	bt.Root = root

	assert.Equal(t,
		[][]interface{}{
			[]interface{}{1},
			[]interface{}{2, 3},
			[]interface{}{4, 5, 6, 7},
		}, bt.BFS2())
}

func TestBinaryTree_Serialize(t *testing.T) {
	bt := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	root.SetLeft(NewBinaryTreeNode(2))
	root.SetRight(NewBinaryTreeNode(3))

	root.Left.SetLeft(NewBinaryTreeNode(4))

	root.Right.SetLeft(NewBinaryTreeNode(6))
	root.Right.SetRight(NewBinaryTreeNode(7))
	bt.Root = root

	assert.Equal(t, "1,2,4,$,$,$,3,6,$,$,7,$", bt.Serialize())
}

func TestDeserialize(t *testing.T) {
	str := []string{"1", "2", "4", "$", "$", "$", "3", "6", "$", "$", "7", "$", "$"}
	//str := "1,2,4,$,$,$,3,6,$,$,7,$,$"
	root := Deserialize(str)
	bt := NewBinaryTree()
	bt.Root = root

	t.Log(bt.Serialize())
}

//todo 递归比非递归版本最后会多一个$
func TestSerializeByRecursion(t *testing.T) {
	bt := NewBinaryTree()
	root := NewBinaryTreeNode(1)
	root.SetLeft(NewBinaryTreeNode(2))
	root.SetRight(NewBinaryTreeNode(3))

	root.Left.SetLeft(NewBinaryTreeNode(4))

	root.Right.SetLeft(NewBinaryTreeNode(6))
	root.Right.SetRight(NewBinaryTreeNode(7))
	bt.Root = root

	assert.Equal(t, "1,2,4,$,$,$,3,6,$,$,7,$,$", strings.Join(SerializeByRecursion(bt.Root), ","))
}
