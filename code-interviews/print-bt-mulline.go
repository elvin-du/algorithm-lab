//把二叉树打印成多行
package code_interviews

import (
	"algorithm-lab/common"
	"fmt"
)
func PrintMulLine(root *common.BinaryTreeNode) []string {
	q := common.NewQueue()
	q.Push(root)
	arrStr := []string{}
	for ; !q.IsEmpty(); {
		str := ""
		size := q.Size()
		for i := 0; i < size; i++ {
			p := q.Pop().(*common.BinaryTreeNode)
			str += fmt.Sprintf("%v", p.Value)
			fmt.Printf("%v",p.Value)

			if !common.IsNil(p.Left) {
				q.Push(p.Left)
			}
			if !common.IsNil(p.Right) {
				q.Push(p.Right)
			}
		}
		fmt.Printf("\n")
		fmt.Println()

		arrStr = append(arrStr, str)
	}

	return arrStr
}
