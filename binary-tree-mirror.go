//剑指offer 二叉树镜像
package algorithm_lab

import "algorithm-lab/common"

func BinaryTreeMirror(tree *common.BinaryTree) {
	binaryTreeMirror(tree.Root)
}

func binaryTreeMirror(node *common.BinaryTreeNode) {
	if common.NilBinaryTreeNode == node {
		return
	}

	tmp := node.Left
	node.SetLeft(node.Right)
	node.SetRight(tmp)

	binaryTreeMirror(node.Left)
	binaryTreeMirror(node.Right)
}
