//剑指offer 合并两个已经排序的列表
package algorithm_lab

func MergeSortedList(l1, l2 *List) *List {
	l1node, l2node := l1.Head, l2.Head

	for ; !IsNil(l1node) && !IsNil(l2node); {
		if l1node.Value.(int) > l2node.Value.(int) {
			l1node = l1node.Next
		} else {
			l2Next := l2node.Next
			l1Next := l1node.Next

			l1node.Next = l2node
			l2node.Next = l1Next

			l2node = l2Next
			l1node = l1Next
		}

		if IsNil(l1node) || IsNil(l2node) || IsNil(l1node.Next) || IsNil(l2node.Next) {
			break
		}
	}

	if IsNil(l1node.Next) {
		if !IsNil(l2node) {
			l1node.Next = l2node
		}
	}

	return l1
}
