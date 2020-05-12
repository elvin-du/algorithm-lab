package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFirstCommonNode(t *testing.T) {
	head1 := NewListNode(1)
	head1Next1 := NewListNode(3)
	head1Next2 := NewListNode(5)
	head1Next3 := NewListNode(7)
	head1Next4 := NewListNode(9)

	head1.Next = head1Next1
	head1Next1.Next = head1Next2
	head1Next2.Next = head1Next3
	head1Next3.Next = head1Next4

	head2 := NewListNode(2)
	head2Next1 := NewListNode(4)
	head2Next2 := NewListNode(6)
	head2.Next = head2Next1
	head2Next1.Next = head2Next2

	common1 := NewListNode(22)
	common2 := NewListNode(33)
	common1.Next = common2

	head2Next2.Next = common1
	head1Next4.Next = common1

	assert.Equal(t, 22, FindFirstCommonNode(head1, head2).Value)

}
