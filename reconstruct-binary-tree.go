//剑指offer 根据前序遍历和中序遍历重新构建二叉树
package algorithm_lab

import "algorithm-lab/common"

func ReconstructBinaryTree(preorder, inorder []int) *common.BinaryTree {
	root := reconstructBinaryTree(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
	bt := common.NewBinaryTree()
	bt.Root = root
	return bt
}

func reconstructBinaryTree(preorder []int, prestart, preend int, inorder []int, instart, inend int) *common.BinaryTreeNode {
	if preend < prestart || instart > inend {
		return nil
	}

	root := common.NewBinaryTreeNode(preorder[prestart])

	for i := instart; i <= inend; i++ {
		if preorder[prestart] == inorder[i] {
			count := i - instart
			root.SetLeft(reconstructBinaryTree(preorder, prestart+1, prestart+count, inorder, instart, i-1))
			root.SetRight(reconstructBinaryTree(preorder, prestart+count+1, preend, inorder, i+1, inend))
		}
	}

	return root
}
