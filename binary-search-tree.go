//二叉搜索树
package algorithm_lab

func CreateBST(data []int) *BinaryTree {
	bt := NewBinaryTree()
	bt.Root = NewBinaryTreeNode(data[0])
	for i := 1; i < len(data); i++ {
		InsertNode(data[i], bt.Root)
	}

	return bt
}

func InsertNode(value int, root *BinaryTreeNode) {
	if IsNil(root) {
		return
	}

	if root.Value.(int) < value {
		if IsNil(root.Right) {
			root.Right = NewBinaryTreeNode(value)
			return
		}

		InsertNode(value, root.Right)
	} else {
		if IsNil(root.Left) {
			root.Left = NewBinaryTreeNode(value)
			return
		}

		InsertNode(value, root.Left)
	}
}

//剑指offer 查找搜索二叉树的第k个节点
func FindKthNode(root *BinaryTreeNode, k *int) *BinaryTreeNode {
	if !IsNil(root) {
		v := FindKthNode(root.Left, k)
		if !IsNil(v) {
			return v
		}

		*k--
		if *k == 0 {
			return root
		}

		v2 := FindKthNode(root.Right, k)
		if !IsNil(v2) {
			return v2
		}
	}

	return nil
}

//剑指offer 查找搜索二叉树的第k个节点
func FindKthNode2(root *BinaryTreeNode, k *int, ret *BinaryTreeNode) {
	if !IsNil(root) {
		FindKthNode2(root.Left, k, ret)

		*k--
		if *k == 0 {
			*ret = *root
		}

		FindKthNode2(root.Right, k, ret)
	}
}
