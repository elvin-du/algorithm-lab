//剑指offer 树的子结构
package algorithm_lab

import "algorithm-lab/common"

func HasSubtree(parent, sub *common.BinaryTree) bool {
	subRoot := sub.Root

	q := common.NewQueue()
	q.Push(parent.Root)
	for ; !q.IsEmpty(); {
		count := q.Size()
		for i := 0; i < count; i++ {
			e := q.Pop().(*common.BinaryTreeNode)
			if !common.IsNil(e.Left) {
				q.Push(e.Left)
			}
			if !common.IsNil(e.Right) {
				q.Push(e.Right)
			}

			if e.Value == subRoot.Value {
				if IsSubTree(e, subRoot) {
					return true
				}
			}
		}
	}

	return false
}

func IsBTSame(a, b *common.BinaryTreeNode) bool {
	if common.IsNil(a) && common.IsNil(b) {
		return true
	}

	if (!common.IsNil(a) && common.IsNil(b)) || (!common.IsNil(b) && common.IsNil(a)) {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return IsBTSame(a.Left, b.Left) && IsBTSame(a.Right, b.Right)
}

func IsSubTree(parent, sub *common.BinaryTreeNode) bool {
	if common.IsNil(parent) && common.IsNil(sub) {
		return true
	}

	if !common.IsNil(parent) && common.IsNil(sub) {
		return true
	}

	if !common.IsNil(sub) && common.IsNil(parent) {
		return false
	}

	if parent.Value != sub.Value {
		return false
	}

	return IsSubTree(parent.Left, sub.Left) && IsSubTree(parent.Right, sub.Right)
}
