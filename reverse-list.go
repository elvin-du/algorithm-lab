//剑指offer 反转链表
package algorithm_lab

import "fmt"

type ListNode_ForReverse struct {
	Next  *ListNode_ForReverse
	Value interface{}
}

func NewListNode_ForReverse(value interface{}) *ListNode_ForReverse {
	return &ListNode_ForReverse{Value: value}
}

func (n *ListNode_ForReverse) String() string {
	str := ""
	for node := n; !IsNil(node); node = node.Next {
		str += fmt.Sprintf("%v", node.Value)
	}
	return str
}

func ReverseList(head *ListNode_ForReverse) *ListNode_ForReverse {
	if IsNil(head) || IsNil(head.Next) {
		return head
	}

	pre, cur := head, head.Next
	head.Next = nil
	for ; !IsNil(cur); {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

func ReverseListByRecursion(head *ListNode_ForReverse) *ListNode_ForReverse {
	if IsNil(head) || IsNil(head.Next) {
		return head
	}

	newList := ReverseListByRecursion(head.Next)

	head.Next.Next = head
	head.Next = nil

	return newList
}
