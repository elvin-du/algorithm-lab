package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := NewListNode_ForReverse(1)
	next2 := NewListNode_ForReverse(2)
	next3 := NewListNode_ForReverse(3)
	next4 := NewListNode_ForReverse(4)

	head.Next = next2
	next2.Next = next3
	next3.Next = next4

	reversedList := ReverseList(head)
	assert.Equal(t, "4321", reversedList.String())
}
