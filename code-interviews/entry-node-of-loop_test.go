package code_interviews

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntryNodeOfLoop(t *testing.T) {
	node1 := common.NewListNode(1)
	node2 := common.NewListNode(2)
	node3 := common.NewListNode(3)
	node4 := common.NewListNode(4)
	node5 := common.NewListNode(5)
	node6 := common.NewListNode(6)

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node4

	assert.Equal(t, 4, EntryNodeOfLoop(node1).Value)
}