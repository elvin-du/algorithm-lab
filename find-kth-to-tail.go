//剑指offer 链表中第k个节点
package algorithm_lab

import "algorithm-lab/common"

//解题思路：双指针，第二个指针比第一个指针后移k个节点
func FindKthToTail(k int, list *common.List) interface{} {
	first := list.Head
	second := list.Head
	i := 0
	for ; i < k-1; i++ {
		if second != nilListNode {
			second = second.Next
		}
	}

	if i != k-1 {
		panic("size of list is less k")
	}

	for ; second.Next != nilListNode; {
		second = second.Next
		first = first.Next
	}

	return first.Value
}
