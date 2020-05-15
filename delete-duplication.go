//剑指offer 删除链表中重复的节点
package algorithm_lab

import "algorithm-lab/common"

var (
	nilListNode *common.ListNode = nil
)

func DeleteDuplication(l *common.List) {
	if l.Head == nilListNode || l.Size == 1 {
		return
	}
	vHead := common.NewListNode(0)
	vHead.Next = l.Head

	pre, cur := vHead, vHead.Next
	for ; cur != nilListNode; {
		if deleteNode(cur) {
			pre.Next = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	l.Head = vHead.Next
}

func deleteNode(node *common.ListNode) bool {
	duplication := false
	for ; node != nilListNode && node.Next != nilListNode; {
		if node.Value != node.Next.Value {
			break
		}

		duplication = true
		node.Next = node.Next.Next
	}

	return duplication
}
