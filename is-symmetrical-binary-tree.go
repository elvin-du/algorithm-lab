//剑指offer 是否是对称二叉树
package algorithm_lab

import "algorithm-lab/common"

func IsSymmetricalBinaryTree(tree *common.BinaryTree) bool {
	return isSymmetricalBinaryTree(tree.Root)
}

func isSymmetricalBinaryTree(node *common.BinaryTreeNode) bool {
	if common.IsNil(node.Left) || common.IsNil(node.Right) {
		if common.IsNil(node.Left) && common.IsNil(node.Right) {
			return true
		}

		return false
	}

	if node.Left.Value != node.Right.Value {
		return false
	}

	return isSymmetricalBinaryTree(node.Left) && isSymmetricalBinaryTree(node.Right)
}
