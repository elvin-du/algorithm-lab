//剑指offer 两个链表的第一个公共节点
package algorithm_lab

import "algorithm-lab/common"

func FindFirstCommonNode(head1, head2 *common.ListNode) *common.ListNode {
	h1, h2 := head1, head2
	for ; !common.IsNil(head1) && !common.IsNil(head2) && head1 != head2; {
		if !common.IsNil(head1.Next) {
			head1 = head1.Next
		} else {
			head1 = h2
		}

		if !common.IsNil(head2.Next) {
			head2 = head2.Next
		} else {
			head2 = h1
		}
	}

	return head2
}
