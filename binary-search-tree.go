//二叉搜索树
package algorithm_lab

import "algorithm-lab/common"

func CreateBST(data []int) *common.BinaryTree {
	bt := common.NewBinaryTree()
	bt.Root = common.NewBinaryTreeNode(data[0])
	for i := 1; i < len(data); i++ {
		InsertNode(data[i], bt.Root)
	}

	return bt
}

func InsertNode(value int, root *common.BinaryTreeNode) {
	if common.IsNil(root) {
		return
	}

	if root.Value.(int) < value {
		if common.IsNil(root.Right) {
			root.Right = common.NewBinaryTreeNode(value)
			return
		}

		InsertNode(value, root.Right)
	} else {
		if common.IsNil(root.Left) {
			root.Left = common.NewBinaryTreeNode(value)
			return
		}

		InsertNode(value, root.Left)
	}
}

//剑指offer 查找搜索二叉树的第k个节点
func FindKthNode(root *common.BinaryTreeNode, k *int) *common.BinaryTreeNode {
	if !common.IsNil(root) {
		v := FindKthNode(root.Left, k)
		if !common.IsNil(v) {
			return v
		}

		*k--
		if *k == 0 {
			return root
		}

		v2 := FindKthNode(root.Right, k)
		if !common.IsNil(v2) {
			return v2
		}
	}

	return nil
}

//剑指offer 查找搜索二叉树的第k个节点
func FindKthNode2(root *common.BinaryTreeNode, k *int, ret *common.BinaryTreeNode) {
	if !common.IsNil(root) {
		FindKthNode2(root.Left, k, ret)

		*k--
		if *k == 0 {
			*ret = *root
		}

		FindKthNode2(root.Right, k, ret)
	}
}
