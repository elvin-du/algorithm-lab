//剑指offer 二叉搜索树转为双向链表
package algorithm_lab

func BSTConvertToList(root *BinaryTreeNode) *BinaryTreeNode {
	var lastNode *BinaryTreeNode = nil

	bstConvert(root, &lastNode)

	for ; !IsNil(lastNode.Left); {
		lastNode = lastNode.Left
	}

	return lastNode
}

//right：表示next指针
//left: 表示pre指针
func bstConvert(root *BinaryTreeNode, lastNode **BinaryTreeNode) {
	if IsNil(root) {
		return
	}

	if !IsNil(root.Left) {
		bstConvert(root.Left, lastNode)
	}

	//指向上一个最后的节点
	root.Left = *lastNode
	if !IsNil(*lastNode) {
		(*lastNode).Right = root
	}
	*lastNode = root

	if !IsNil(root.Right) {
		bstConvert(root.Right, lastNode)
	}
}
