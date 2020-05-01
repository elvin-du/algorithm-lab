//剑指offer 中序遍历的下一个节点
package algorithm_lab

func GetNextNode(node *BinaryTreeNode) *BinaryTreeNode {
	//有右子树
	if node.Right != NilBinaryTreeNode {
		for n := node.Right; n != NilBinaryTreeNode; n = n.Left {
			if n.Left == NilBinaryTreeNode {
				return n
			}
		}
	} else if node.Right == NilBinaryTreeNode &&
		node.Parent != NilBinaryTreeNode &&
		node.Parent.Left == node {
		return node.Parent
	} else if node.Right == NilBinaryTreeNode &&
		node.Parent != NilBinaryTreeNode &&
		node.Parent.Right == node {
		for pnode := node.Parent; nil != pnode; pnode = pnode.Parent {

			if NilBinaryTreeNode != pnode.Parent && pnode.Parent.Left == pnode {
				return pnode.Parent
			}
		}
	}

	return nil
}
