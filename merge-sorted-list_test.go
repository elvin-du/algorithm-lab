package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSortedList(t *testing.T) {
	l1 := NewList()
	l2 := NewList()

	l1.Head = NewListNode(1)
	l1.Head.Next = NewListNode(3)
	l1.Head.Next.Next = NewListNode(5)
	l1.Head.Next.Next.Next = NewListNode(7)

	l2.Head = NewListNode(2)
	l2.Head.Next = NewListNode(4)
	l2.Head.Next.Next = NewListNode(6)
	l2.Head.Next.Next.Next = NewListNode(8)

	l3 := MergeSortedList(l1, l2)
	assert.Equal(t, "12345678", l3.String())
}

func TestMergeSortedList2(t *testing.T) {
	l1 := NewList()
	l2 := NewList()

	l1.Head = NewListNode(1)
	l1.Head.Next = NewListNode(3)
	l1.Head.Next.Next = NewListNode(5)
	l1.Head.Next.Next.Next = NewListNode(7)

	l3 := MergeSortedList(l1, l2)
	assert.Equal(t, "1357", l3.String())
}

func TestMergeSortedList3(t *testing.T) {
	l1 := NewList()
	l2 := NewList()

	l1.Head = NewListNode(1)
	l1.Head.Next = NewListNode(3)
	l1.Head.Next.Next = NewListNode(5)
	l1.Head.Next.Next.Next = NewListNode(7)

	l2.Head = NewListNode(9)

	l3 := MergeSortedList(l1, l2)
	assert.Equal(t, "13579", l3.String())
}
