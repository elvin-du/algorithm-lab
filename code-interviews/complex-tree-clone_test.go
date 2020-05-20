package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCloneComplexTree(t *testing.T) {
	head := NewComplexListNode(1)
	next2 := NewComplexListNode(2)
	next3 := NewComplexListNode(3)
	next4 := NewComplexListNode(4)
	next5 := NewComplexListNode(5)

	head.Next = next2
	next2.Next = next3
	next3.Next = next4
	next4.Next = next5

	next2.Random = next5
	next4.Random = next2

	newHead := CloneComplexTree(head)
	assert.Equal(t, 5, newHead.Next.Random.Value)
	assert.Equal(t, 2, newHead.Next.Next.Next.Random.Value)
}