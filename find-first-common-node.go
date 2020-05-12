//剑指offer 两个链表的第一个公共节点
package algorithm_lab

func FindFirstCommonNode(head1, head2 *ListNode) *ListNode {
	h1, h2 := head1, head2
	for ; !IsNil(head1) && !IsNil(head2) && head1 != head2; {
		if !IsNil(head1.Next) {
			head1 = head1.Next
		} else {
			head1 = h2
		}

		if !IsNil(head2.Next) {
			head2 = head2.Next
		} else {
			head2 = h1
		}
	}

	return head2
}
