//中序遍历的下一个节点
package code_interviews

import "algorithm-lab/common"

func NextNodeInorder(node *common.BinaryTreeNode) *common.BinaryTreeNode {
	if !common.IsNil(node.Right) {
		n := node.Right
		for ; !common.IsNil(n.Left); {
			n = n.Left
		}
		return n
	} else if !common.IsNil(node.Parent) && node.Parent.Left == node {
		return node.Parent
	} else if !common.IsNil(node.Parent) && node.Parent.Right == node {
		n := node
		for ; !common.IsNil(n.Parent) && n.Parent.Right == n; {
			n = n.Parent
		}

		if !common.IsNil(n.Parent) && n.Parent.Left == n {
			return n.Parent
		}
	}

	return nil
}
