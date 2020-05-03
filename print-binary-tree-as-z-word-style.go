//剑指offer 按照Z字型打印二叉树
package algorithm_lab

import (
	"fmt"
)

func PrintBinaryTreeAsZWordStyle(tree *BinaryTree) string {
	//奇数层栈
	stack1 := NewStack(100)
	//偶数层栈
	stack2 := NewStack(100)
	str := ""
	stack1.Push(tree.Root)
	for {
		if stack1.Size() == 0 {
			break
		}
		for ; stack1.Size() != 0; {
			e := stack1.Pop().(*BinaryTreeNode)
			fmt.Print(e.Value, " ")
			str += fmt.Sprintf("%v ", e.Value)

			if e.Left != NilBinaryTreeNode {
				stack2.Push(e.Left)
			}
			if e.Right != NilBinaryTreeNode {
				stack2.Push(e.Right)
			}
		}
		fmt.Println()

		for ; stack2.Size() != 0; {
			e := stack2.Pop().(*BinaryTreeNode)
			fmt.Print(e.Value, " ")
			str += fmt.Sprintf("%v ", e.Value)
			if e.Right != NilBinaryTreeNode {
				stack1.Push(e.Right)
			}
			if e.Left != NilBinaryTreeNode {
				stack1.Push(e.Left)
			}
		}
		fmt.Println()
	}

	return str
}
