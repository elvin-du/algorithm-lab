package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindKthToTail(t *testing.T) {
	list := NewList()
	head := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node4 := NewListNode(4)
	node5 := NewListNode(5)
	node6 := NewListNode(6)
	node7 := NewListNode(7)
	node8 := NewListNode(8)
	node9 := NewListNode(9)

	list.Head = head
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node9

	assert.Equal(t, 7, FindKthToTail(3, list).(int))
}
