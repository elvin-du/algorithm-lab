package code_interviews

func FindContinuousSequence(n int) [][]int {
	pre, cur := 1, 2
	ret := [][]int{}
	for ; pre < cur && cur < n/2; {
		if Sum(pre, cur) > n {
			pre++
		} else if Sum(pre, cur) < n {
			cur++
		} else {
			tmp := []int{}
			for i := pre; i <= cur; i++ {
				tmp = append(tmp, i)
			}
			ret = append(ret, tmp)
			pre++
		}
	}

	return ret
}

func Sum(l, r int) int {
	sum := 0
	for ; l <= r; l++ {
		sum += l
	}
	return sum
}
