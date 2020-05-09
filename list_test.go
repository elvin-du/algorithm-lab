package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	list := NewList()
	assert.Equal(t, -1, list.Find(1))
	list.PushBack(1)
	assert.Equal(t, 1, list.RemoveByIndex(0))
	list.PushFront(1)
	assert.Equal(t, 0, list.Find(1))
	list.PushBack(2)
	list.PushBack(3)

	assert.Equal(t, 3, list.Size)
	assert.Equal(t, 2, list.Find(3))
	assert.Equal(t, 1, list.RemoveByIndex(0))
	assert.Equal(t, 3, list.RemoveByIndex(1))
	assert.Equal(t, 2, list.RemoveByIndex(0))

	list2 := NewList()
	list2.PushBack(1)
	list2.PushBack(2)
	list2.PushBack(3)
	list2.PushFront(4)
	list2.PushFront(5)
	list2.PushFront(6)
	assert.Equal(t, 6, list2.Get(0))
	assert.Equal(t, 3, list2.Get(5))
}

func TestList_String(t *testing.T) {
	l := NewList()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushFront(4)
	assert.Equal(t, "4321", l.String())
}

func TestList_RemoveByIndex(t *testing.T) {
	list2 := NewList()
	list2.PushBack(1)
	list2.PushBack(2)
	list2.PushBack(3)
	assert.Equal(t, 3, list2.Size)
	assert.Equal(t, 3, list2.RemoveByIndex(list2.Size-1))
	assert.Equal(t, 2, list2.Size)
	assert.Equal(t, 2, list2.RemoveByIndex(list2.Size-1))
	assert.Equal(t, 1, list2.Size)
	assert.Equal(t, 1, list2.RemoveByIndex(list2.Size-1))
}
