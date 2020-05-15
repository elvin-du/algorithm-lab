//剑指offer 从尾到头打印链表
package code_interviews

import (
	"algorithm-lab/common"
	"fmt"
)

func PrintListFromTailToHead(head *common.ListNode) {
	if common.IsNil(head) {
		return
	}
	PrintListFromTailToHead(head.Next)
	fmt.Println(head.Value)
}
