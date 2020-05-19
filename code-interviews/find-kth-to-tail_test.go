package code_interviews

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindKthToTail(t *testing.T) {
	list := common.NewList()
	head := common.NewListNode(1)
	node2 := common.NewListNode(2)
	node3 := common.NewListNode(3)
	node4 := common.NewListNode(4)
	node5 := common.NewListNode(5)
	node6 := common.NewListNode(6)
	node7 := common.NewListNode(7)
	node8 := common.NewListNode(8)
	node9 := common.NewListNode(9)

	list.Head = head
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node9

	assert.Equal(t, 7, FindKthToTail(list.Head, 3))
}
