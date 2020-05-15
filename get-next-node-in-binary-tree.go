//剑指offer 中序遍历的下一个节点
package algorithm_lab

import "algorithm-lab/common"

func GetNextNode(node *common.BinaryTreeNode) *common.BinaryTreeNode {
	//有右子树
	if node.Right != common.NilBinaryTreeNode {
		for n := node.Right; n != common.NilBinaryTreeNode; n = n.Left {
			if n.Left == common.NilBinaryTreeNode {
				return n
			}
		}
	} else if node.Right == common.NilBinaryTreeNode &&
		node.Parent != common.NilBinaryTreeNode &&
		node.Parent.Left == node {
		return node.Parent
	} else if node.Right == common.NilBinaryTreeNode &&
		node.Parent != common.NilBinaryTreeNode &&
		node.Parent.Right == node {
		for pnode := node.Parent; nil != pnode; pnode = pnode.Parent {

			if common.NilBinaryTreeNode != pnode.Parent && pnode.Parent.Left == pnode {
				return pnode.Parent
			}
		}
	}

	return nil
}
