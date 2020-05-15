package algorithm_lab

import (
	"algorithm-lab/common"
	"testing"
)

func TestPrintListFromTailToHead(t *testing.T) {
	l := common.NewList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	PrintListFromTailToHead(l)
}
