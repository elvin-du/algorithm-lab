package code_interviews

import "algorithm-lab/common"

type ComplexListNode struct {
	Next   *ComplexListNode
	Random *ComplexListNode
	Value  int
}

func NewComplexListNode(v int) *ComplexListNode {
	return &ComplexListNode{Value: v}
}

func CloneComplexTree(root *ComplexListNode) *ComplexListNode {
	cloneNext(root)
	recalculatePointer(root)
	return split(root)
}

func split(root *ComplexListNode) *ComplexListNode {
	newRoot := root.Next
	pre, cur := root, root.Next
	for ; !common.IsNil(cur); {
		pre.Next = pre.Next.Next
		pre = pre.Next
		if common.IsNil(cur.Next) {
			break
		}
		cur.Next = cur.Next.Next
		cur = cur.Next
	}

	return newRoot
}

func recalculatePointer(root *ComplexListNode) {
	root = root.Next
	for ; !common.IsNil(root)&& !common.IsNil(root.Next); {
		if !common.IsNil(root.Random){
			root.Random = root.Random.Next
		}
		if !common.IsNil(root.Next) {
			root = root.Next.Next
		}
	}
}

func cloneNext(root *ComplexListNode) {
	for ; !common.IsNil(root); {
		next := NewComplexListNode(root.Value)
		next.Random = root.Random
		next.Next = root.Next

		root.Next = next

		root = root.Next.Next
	}
}
