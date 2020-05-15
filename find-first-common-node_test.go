package algorithm_lab

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFirstCommonNode(t *testing.T) {
	head1 := common.NewListNode(1)
	head1Next1 := common.NewListNode(3)
	head1Next2 := common.NewListNode(5)
	head1Next3 := common.NewListNode(7)
	head1Next4 := common.NewListNode(9)

	head1.Next = head1Next1
	head1Next1.Next = head1Next2
	head1Next2.Next = head1Next3
	head1Next3.Next = head1Next4

	head2 := common.NewListNode(2)
	head2Next1 := common.NewListNode(4)
	head2Next2 := common.NewListNode(6)
	head2.Next = head2Next1
	head2Next1.Next = head2Next2

	common1 := common.NewListNode(22)
	common2 := common.NewListNode(33)
	common1.Next = common2

	head2Next2.Next = common1
	head1Next4.Next = common1

	assert.Equal(t, 22, FindFirstCommonNode(head1, head2).Value)

}
