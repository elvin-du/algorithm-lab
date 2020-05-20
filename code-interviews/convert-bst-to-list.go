package code_interviews

import "algorithm-lab/common"

var (
	pre  *common.BinaryTreeNode
	head *common.BinaryTreeNode
)

func ConvertBSTToList(root *common.BinaryTreeNode) {
	if common.IsNil(root) {
		return
	}

	ConvertBSTToList(root.Left)
	if common.IsNil(head) {
		head = root
	}
	if !common.IsNil(pre) {
		pre.Right = root
	}
	root.Left = pre
	pre = root
	ConvertBSTToList(root.Right)
}

