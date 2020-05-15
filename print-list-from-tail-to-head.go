//剑指offer 从头到尾打印链表
package algorithm_lab

import (
	"algorithm-lab/common"
	"log"
)

func printListFromTailToHead(l *common.ListNode) {
	if l == nil {
		return
	}
	printListFromTailToHead(l.Next)
	log.Println(l.Value)
}

func PrintListFromTailToHead(l *common.List) {
	printListFromTailToHead(l.Head)
}
