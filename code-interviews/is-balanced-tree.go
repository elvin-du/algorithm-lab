package code_interviews

import (
	"algorithm-lab/common"
	"math"
)

var isBalanced = true

func IsBalancedTree(root *common.BinaryTreeNode) bool {
	height(root)
	return isBalanced
}

func height(root *common.BinaryTreeNode) int {
	if common.IsNil(root) || !isBalanced {
		return 0
	}

	left := height(root.Left)
	right := height(root.Right)
	if math.Abs(float64(left-right)) > 1 {
		isBalanced = false
	}

	return 1 + int(math.Max(float64(left), float64(right)))
}
