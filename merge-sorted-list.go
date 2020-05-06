//剑指offer 合并两个已经排序的列表
package algorithm_lab

import "math"

func MergeSortedList(l1, l2 *List) *List {
	l1node, l2node := l1.Head, l2.Head
	prel1 := NewListNode(math.MinInt64)
	prel1.Next = l1node

	for ; !IsNil(l2node); {
		for ; !IsNil(l1node); {
			if l1node.Value.(int) < l2node.Value.(int) {
				prel1 = l1node
				l1node = l1node.Next
			} else {
				tmpL2Next := l2node.Next

				prel1.Next = l2node
				prel1.Next.Next = l1node
				prel1 = prel1.Next

				l2node = tmpL2Next

				break
			}
		}

		if IsNil(l1node) {
			break
		}
	}

	if !IsNil(l2node) {
		prel1.Next = l2node
	}

	return l1
}
