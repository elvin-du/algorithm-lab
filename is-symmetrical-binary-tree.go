//剑指offer 是否是对称二叉树
package algorithm_lab

import (
	"algorithm-lab/common"
)

func IsSymmetricalBinaryTree(tree *common.BinaryTree) bool {
	if common.IsNil(tree.Root) {
		return true
	}

	return isSymmetricalBinaryTree(tree.Root.Left, tree.Root.Right)
}

func isSymmetricalBinaryTree(left *common.BinaryTreeNode, right *common.BinaryTreeNode) bool {
	if common.IsNil(left) && common.IsNil(right) {
		return true
	}

	if common.IsNil(left) || common.IsNil(right) {
		return false
	}

	if left.Value != right.Value {
		return false
	}

	return isSymmetricalBinaryTree(left.Left, right.Right) && isSymmetricalBinaryTree(left.Right, right.Left)
}
