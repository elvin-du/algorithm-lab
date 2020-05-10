//剑指offer 树的子结构
package algorithm_lab

func HasSubtree(parent, sub *BinaryTree) bool {
	subRoot := sub.Root

	q := NewQueue()
	q.Push(parent.Root)
	for ; !q.IsEmpty(); {
		count := q.Size()
		for i := 0; i < count; i++ {
			e := q.Pop().(*BinaryTreeNode)
			if !IsNil(e.Left) {
				q.Push(e.Left)
			}
			if !IsNil(e.Right) {
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

func IsBTSame(a, b *BinaryTreeNode) bool {
	if IsNil(a) && IsNil(b) {
		return true
	}

	if (!IsNil(a) && IsNil(b)) || (!IsNil(b) && IsNil(a)) {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return IsBTSame(a.Left, b.Left) && IsBTSame(a.Right, b.Right)
}

func IsSubTree(parent, sub *BinaryTreeNode) bool {
	if IsNil(parent) && IsNil(sub) {
		return true
	}

	if !IsNil(parent) && IsNil(sub) {
		return true
	}

	if !IsNil(sub) && IsNil(parent) {
		return false
	}

	if parent.Value != sub.Value {
		return false
	}

	return IsSubTree(parent.Left, sub.Left) && IsSubTree(parent.Right, sub.Right)
}
