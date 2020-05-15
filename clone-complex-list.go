//剑指offer 复制复杂链表
package algorithm_lab

import "algorithm-lab/common"

type ComplexListNode struct {
	Next   *ComplexListNode
	Random *ComplexListNode
	Value  int
}

func NewComplexListNode(v int) *ComplexListNode {
	return &ComplexListNode{Value: v}
}

func CloneComplexList(head *ComplexListNode) *ComplexListNode {
	cloneComplexListNode(head)
	ConnectSiblingNodes(head)
	return ReconnectNodes(head)
}

func cloneComplexListNode(head *ComplexListNode) {
	for node := head; !common.IsNil(node); {
		newNode := NewComplexListNode(node.Value)
		newNode.Random = node.Random

		next := node.Next
		node.Next = newNode
		newNode.Next = next

		node = node.Next.Next
	}
}

func ConnectSiblingNodes(head *ComplexListNode) {
	i := 1
	for node := head; !common.IsNil(node); {
		if i%2 == 0 {
			if !common.IsNil(node.Random) {
				node.Random = node.Random.Next
			}
		}
		node = node.Next
		i++
	}
}

func ReconnectNodes(head *ComplexListNode) *ComplexListNode {
	newHead := head.Next

	for node := newHead; !common.IsNil(node) && !common.IsNil(node.Next); {
		if !common.IsNil(node.Next) {
			node.Next = node.Next.Next
			node = node.Next
		}
	}

	return newHead
}
