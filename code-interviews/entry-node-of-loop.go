//剑指offer 链表中环的入口节点
package code_interviews

import "algorithm-lab/common"

//假定已经存在环
func EntryNodeOfLoop(head *common.ListNode) *common.ListNode {
	slow, fast := head.Next, head.Next.Next
	for ; fast != slow; {
		fast = fast.Next.Next
		slow = slow.Next
	}

	fast = head
	for ; slow != fast; {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
