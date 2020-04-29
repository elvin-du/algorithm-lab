package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntryNodeOfLoop(t *testing.T) {
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node4 := NewListNode(4)
	node5 := NewListNode(5)
	node6 := NewListNode(6)

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node4

	assert.Equal(t, 4, EntryNodeOfLoop(node1).Value)

}

//没有环
func TestEntryNodeOfLoop2(t *testing.T) {
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node4 := NewListNode(4)
	node5 := NewListNode(5)
	node6 := NewListNode(6)

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	assert.Nil(t, EntryNodeOfLoop(node1))
}
