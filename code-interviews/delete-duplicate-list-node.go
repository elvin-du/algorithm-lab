//删除链表中重复的节点
package code_interviews

import "algorithm-lab/common"

func DeleteDuplicationNode(l *common.List) {
	vHead := common.NewListNode(0)
	vHead.Next = l.Head
	pre, cur := vHead, vHead.Next
	for ; !common.IsNil(cur); {
		if DeleteNode(cur) {
			pre.Next = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	l.Head = vHead.Next
}

func DeleteNode(node *common.ListNode) bool {
	duplication := false
	for ; !common.IsNil(node) && !common.IsNil(node.Next); {
		if node.Value != node.Next.Value {
			break
		}
		duplication = true
		node.Next = node.Next.Next
	}

	return duplication
}
