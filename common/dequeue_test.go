package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDequeue(t *testing.T) {
	dq := NewDequeue()
	dq.PushBack(5)
	dq.PushBack(6)
	dq.PushFront(4)
	dq.PushFront(3)

	assert.Equal(t, "3456", dq.list.String())

	assert.Equal(t, 6, dq.PeekBack())
	assert.Equal(t, 3, dq.PeekFront())

	assert.Equal(t, 3, dq.PollFront())
	assert.Equal(t, 6, dq.PollBack())
}
