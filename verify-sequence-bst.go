//剑指offer 二叉搜索树的后序遍历序列
package algorithm_lab

func VerifySequenceOfBST(sequence []int) bool {
	return verifySequenceOfBST(sequence, 0, len(sequence)-1)
}

func verifySequenceOfBST(sequence []int, start, end int) bool {
	if start >= end {
		return true
	}

	root := sequence[end]
	index := rootIndex(sequence, start, end, root)
	if -1 == index {
		return false
	}

	return verifySequenceOfBST(sequence, start, index-1) && verifySequenceOfBST(sequence, index, end-1)
}

func rootIndex(sequence []int, start int, end int, root int) int {
	index := -1
	for ; start <= end; {
		if sequence[start] >= root {
			index = start
			break
		}
		start++
	}

	for start++; start < end; {
		if sequence[start] < root {
			return -1
		}

		start++
	}

	return index
}
