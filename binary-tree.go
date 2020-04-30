package algorithm_lab

import (
	"fmt"
	"strings"
)

var NilBinaryTreeNode *BinaryTreeNode = nil

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) PreOrder() string {
	return preOrder(bt.Root)
}
func (bt *BinaryTree) InOrder() string {
	return inOrder(bt.Root)
}

func inOrder(node *BinaryTreeNode) string {
	if NilBinaryTreeNode == node {
		return ""
	}

	str := ""
	str += inOrder(node.Left)
	str += fmt.Sprintf("%v", node.Value)
	str += inOrder(node.Right)
	return str
}

//宽度优先遍历
func (bt *BinaryTree) BFS() string {
	if bt.Root == NilBinaryTreeNode {
		return ""
	}

	queue := NewQueue()
	queue.Push(bt.Root)

	strArry := []string{}
	for ; !queue.IsEmpty(); {
		n := queue.Pop().(*BinaryTreeNode)
		strArry = append(strArry, fmt.Sprintf("%v", n.Value))
		if n.Left != NilBinaryTreeNode {
			queue.Push(n.Left)
		}
		if n.Right != NilBinaryTreeNode {
			queue.Push(n.Right)
		}
	}

	return strings.Join(strArry, ",")
}

//默认是宽度优先遍历
func (bt *BinaryTree) String() string {
	return bt.BFS()
}

func preOrder(node *BinaryTreeNode) string {
	if NilBinaryTreeNode == node {
		return ""
	}
	str := fmt.Sprintf("%v", node.Value)

	str += preOrder(node.Left)
	str += preOrder(node.Right)

	return str
}

type BinaryTreeNode struct {
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
	Parent *BinaryTreeNode
	Value  interface{}
}

func NewBinaryTreeNode(value interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{Value: value}
}

func (node *BinaryTreeNode) SetLeft(n *BinaryTreeNode) {
	node.Left = n
	if NilBinaryTreeNode != n {
		n.Parent = node
	}
}

func (node *BinaryTreeNode) SetRight(n *BinaryTreeNode) {
	node.Right = n
	if NilBinaryTreeNode != n {
		n.Parent = node
	}
}
