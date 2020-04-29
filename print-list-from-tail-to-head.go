//剑指offer 从头到尾打印链表
package algorithm_lab

import "log"

func printListFromTailToHead(l *ListNode) {
	if l == nil {
		return
	}
	printListFromTailToHead(l.Next)
	log.Println(l.Value)
}

func PrintListFromTailToHead(l *List) {
	printListFromTailToHead(l.Head)
}
