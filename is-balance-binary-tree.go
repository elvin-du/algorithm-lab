//剑指offer 判断是否是平衡二叉树
package algorithm_lab

import (
	"algorithm-lab/common"
	"math"
)

var isBalanced = true

func IsBalancedBinaryTree(root *common.BinaryTreeNode) bool {
	height(root)
	return isBalanced
}

func height(root *common.BinaryTreeNode) int {
	if common.IsNil(root) {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)

	if math.Abs(float64(left-right)) > 1 {
		isBalanced = false
	}

	return 1 + Max(left, right)
}
