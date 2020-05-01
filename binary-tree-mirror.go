//剑指offer 二叉树镜像
package algorithm_lab

func BinaryTreeMirror(tree *BinaryTree) {
	binaryTreeMirror(tree.Root)
}

func binaryTreeMirror(node *BinaryTreeNode) {
	if NilBinaryTreeNode == node {
		return
	}

	tmp := node.Left
	node.SetLeft(node.Right)
	node.SetRight(tmp)

	binaryTreeMirror(node.Left)
	binaryTreeMirror(node.Right)
}
