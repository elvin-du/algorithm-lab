//剑指offer 链表中环的入口节点
package algorithm_lab

import (
	"algorithm-lab/common"
	"reflect"
)

func EntryNodeOfLoop(node *common.ListNode) *common.ListNode {
	head := node
	fast, slow := node.Next.Next, node.Next
	for ; !reflect.ValueOf(fast).IsNil() && slow != fast; {
		fast = fast.Next
		if reflect.ValueOf(fast).IsNil() {
			return nil
		}

		slow = slow.Next
		fast = fast.Next
	}

	if reflect.ValueOf(fast).IsNil() {
		return nil
	}

	for ; head != slow; head, slow = head.Next, slow.Next {
	}
	return slow
}
