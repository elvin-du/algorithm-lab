package code_interviews

func IsBST(data []int) bool {
	if len(data) <= 1 {
		return true
	}

	i := Index(data, data[len(data)-1])
	if -1 == i {
		return false
	}

	if i == len(data) {
		left, right := data[:len(data)-1], []int{}
		return IsBST(left) && IsBST(right)
	} else if 0 == i {
		left, right := []int{}, data[:len(data)-1]
		return IsBST(left) && IsBST(right)
	}

	left, right := data[0:i], data[i:len(data)-1]
	return IsBST(left) && IsBST(right)
}

func Index(data []int, k int) int {
	i := 0
	for ; i < len(data); i++ {
		if data[i] > k {
			break
		}
	}
	index := i
	for ; i < len(data); i++ {
		if data[i] < k {
			return -1
		}
	}

	return index
}
