//剑指offer 合并两个已经排序的列表
package algorithm_lab

import (
	"algorithm-lab/common"
	"math"
)

func MergeSortedList(l1, l2 *common.List) *common.List {
	l1node, l2node := l1.Head, l2.Head
	prel1 := common.NewListNode(math.MinInt64)
	prel1.Next = l1node

	for ; !common.IsNil(l2node); {
		for ; !common.IsNil(l1node); {
			if l1node.Value.(int) < l2node.Value.(int) {
				prel1 = l1node
				l1node = l1node.Next
			} else {
				tmpL2Next := l2node.Next

				prel1.Next = l2node
				l2node.Next = l1node

				prel1 = prel1.Next

				l2node = tmpL2Next

				break
			}
		}

		if common.IsNil(l1node) {
			break
		}
	}

	if !common.IsNil(l2node) {
		prel1.Next = l2node
	}

	return l1
}
