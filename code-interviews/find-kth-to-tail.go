package code_interviews

import "algorithm-lab/common"

func FindKthToTail(head *common.ListNode, k int) int {
	slow, fast := head, head
	for i := 0; i < k-1; i++ {
		if !common.IsNil(fast.Next) {
			fast = fast.Next
		} else {
			return -1
		}
	}

	for ; !common.IsNil(fast.Next); {
		slow = slow.Next
		fast = fast.Next
	}

	return slow.Value.(int)
}
