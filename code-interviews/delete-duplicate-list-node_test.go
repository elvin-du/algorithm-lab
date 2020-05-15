package code_interviews

import (
	"algorithm-lab/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicationNode(t *testing.T) {
	l := common.NewList()

	l.PushBack(1)
	l.PushBack(1)
	l.PushBack(2)
	DeleteDuplicationNode(l)
	assert.Equal(t, 2, l.Get(0))

	l.PushBack(3)
	l.PushBack(3)
	l.PushBack(4)
	assert.Equal(t, "2334", l.String())
	DeleteDuplicationNode(l)
	assert.Equal(t, "24", l.String())
	l2 := common.NewList()
	assert.Equal(t, "", l2.String())
	l2.PushBack(2)
	l2.PushBack(2)
	DeleteDuplicationNode(l2)
	assert.Equal(t, "", l2.String())
}