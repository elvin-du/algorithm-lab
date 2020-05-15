//剑指offer 二叉树的深度
package algorithm_lab

import "algorithm-lab/common"

func TreeDepth(root *common.BinaryTreeNode) int {
	if common.IsNil(root) {
		return 0
	}

	return 1 + Max(TreeDepth(root.Left), TreeDepth(root.Right))
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
