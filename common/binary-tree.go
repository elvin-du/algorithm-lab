package common

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

//剑指offer 按行打印每层数据
func (bt *BinaryTree) BFS2() [][]interface{} {
	ret := [][]interface{}{}
	queue := NewQueue()
	queue.Push(bt.Root)

	for ; !queue.IsEmpty(); {
		cnt := queue.Size()
		layer := []interface{}{}
		for i := 0; i < cnt; i++ {
			e := queue.Pop().(*BinaryTreeNode)
			if IsNil(e) {
				continue
			}
			layer = append(layer, e.Value)

			queue.Push(e.Left)
			queue.Push(e.Right)
		}
		if len(layer) != 0 {
			ret = append(ret, layer)
		}
	}

	return ret
}

func SerializeByRecursion(root *BinaryTreeNode) []string {
	strArr := []string{}
	if IsNil(root) {
		strArr = append(strArr, "$")
		return strArr
	}

	strArr = append(strArr, fmt.Sprintf("%v", root.Value))

	strArr = append(strArr, SerializeByRecursion(root.Left)...)
	strArr = append(strArr, SerializeByRecursion(root.Right)...)

	return strArr
}

//非递归版本的前序遍历
func (bt *BinaryTree) Serialize() string {
	stack := NewStack(1000)
	node := bt.Root
	ret := []string{}
	for {
		if !IsNil(node) {
			ret = append(ret, fmt.Sprintf("%v", node.Value))
			stack.Push(node)
			node = node.Left
		} else {
			ret = append(ret, "$")
			e := stack.Pop().(*BinaryTreeNode)
			node = e.Right
		}
		if stack.Size() == 0 && IsNil(node) {
			break
		}
	}

	return strings.Join(ret, ",")
}

var index = 0

func Deserialize(str []string) *BinaryTreeNode {
	if len(str) == 0 {
		return nil
	}
	if index >= len(str) {
		return nil
	}

	if str[index] == "$" {
		index++
		return nil
	}

	root := NewBinaryTreeNode(str[index])
	index++
	root.Left = Deserialize(str)
	root.Right = Deserialize(str)

	return root
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
