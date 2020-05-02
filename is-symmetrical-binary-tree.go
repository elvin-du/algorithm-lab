//剑指offer 是否是对称二叉树
package algorithm_lab

func IsSymmetricalBinaryTree(tree *BinaryTree) bool {
	return isSymmetricalBinaryTree(tree.Root)
}

func isSymmetricalBinaryTree(node *BinaryTreeNode) bool {
	if IsNil(node.Left) || IsNil(node.Right) {
		if IsNil(node.Left) && IsNil(node.Right) {
			return true
		}

		return false
	}

	if node.Left.Value != node.Right.Value {
		return false
	}

	return isSymmetricalBinaryTree(node.Left) && isSymmetricalBinaryTree(node.Right)
}
