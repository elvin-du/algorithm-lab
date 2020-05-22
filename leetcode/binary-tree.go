package leetcode

import (
	"algorithm-lab/common"
	"fmt"
)

func inorder(root *common.BinaryTreeNode) {
	stack := common.NewStack(1000)
	for ; stack.Size() != 0 || nil != root; {
		for ; nil != root; {
			stack.Push(root)
			root = root.Left
		}
		v := stack.Pop().(*common.BinaryTreeNode)
		fmt.Println(v.Value)
		root = v.Right
	}
}
func preorder(root *common.BinaryTreeNode) {
	stack := common.NewStack(1000)
	for ; stack.Size() != 0 || nil != root; {
		for ; nil != root; {
			fmt.Println(root.Value)

			stack.Push(root)
			root = root.Left
		}
		v := stack.Pop().(*common.BinaryTreeNode)
		root = v.Right
	}
}

func postOrder(root *common.BinaryTreeNode) {
	output, stack := common.NewStack(100), common.NewStack(100)
	for ; stack.Size() != 0 || nil != root; {
		if nil != root {
			stack.Push(root)
			output.Push(root)
			root = root.Right
		} else {
			v := stack.Pop().(*common.BinaryTreeNode)
			root = v.Left
		}
	}

	for ; output.Size() != 0; {
		v := output.Pop().(*common.BinaryTreeNode)
		fmt.Println(v.Value)
	}
}
