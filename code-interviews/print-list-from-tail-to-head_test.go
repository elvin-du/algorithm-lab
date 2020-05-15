package code_interviews

import (
	"algorithm-lab/common"
	"testing"
)

func TestPrintListFromTailToHead(t *testing.T) {
	head := common.NewListNode(1)
	node2 := common.NewListNode(2)
	node3 := common.NewListNode(3)
	node4 := common.NewListNode(4)

	head.Next = node2
	node2.Next = node3
	node3.Next = node4

	PrintListFromTailToHead(head)
}
