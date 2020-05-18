//z字型打印二叉树
package code_interviews

import (
	"algorithm-lab/common"
	"fmt"
)

func ZWordPrint(root *common.BinaryTreeNode) string {
	s1, s2 := common.NewStack(100), common.NewStack(100)
	s1.Push(root)
	str := ""
	for {
		if s1.Size() == 0 {
			break
		}
		for ; s1.Size() != 0; {
			p := s1.Pop().(*common.BinaryTreeNode)
			fmt.Printf("%v", p.Value)
			str += fmt.Sprintf("%v ", p.Value)
			if !common.IsNil(p.Left) {
				s2.Push(p.Left)
			}

			if !common.IsNil(p.Right) {
				s2.Push(p.Right)
			}
		}

		for ; s2.Size() != 0; {
			p := s2.Pop().(*common.BinaryTreeNode)
			fmt.Printf("%v", p.Value)
			str += fmt.Sprintf("%v ", p.Value)

			if !common.IsNil(p.Right) {
				s1.Push(p.Right)
			}
			if !common.IsNil(p.Left) {
				s1.Push(p.Left)
			}
		}

		fmt.Println()
	}

	return str
}
