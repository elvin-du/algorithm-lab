package algorithm_lab

import "testing"

func TestPrintListFromTailToHead(t *testing.T) {
	l := NewList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	PrintListFromTailToHead(l)
}
