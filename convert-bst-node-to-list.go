//剑指offer 二叉搜索树转为双向链表
package algorithm_lab

import "algorithm-lab/common"

func BSTConvertToList(root *common.BinaryTreeNode) *common.BinaryTreeNode {
	var lastNode *common.BinaryTreeNode = nil

	bstConvert(root, &lastNode)

	for ; !common.IsNil(lastNode.Left); {
		lastNode = lastNode.Left
	}

	return lastNode
}

//right：表示next指针
//left: 表示pre指针
func bstConvert(root *common.BinaryTreeNode, lastNode **common.BinaryTreeNode) {
	if common.IsNil(root) {
		return
	}

	if !common.IsNil(root.Left) {
		bstConvert(root.Left, lastNode)
	}

	//指向上一个最后的节点
	root.Left = *lastNode
	if !common.IsNil(*lastNode) {
		(*lastNode).Right = root
	}
	*lastNode = root

	if !common.IsNil(root.Right) {
		bstConvert(root.Right, lastNode)
	}
}
