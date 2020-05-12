//剑指offer 数组中出现次数超过一半的数
package algorithm_lab

//假设众数一定存在
func MoreThanHalfNumber(data []int) int {
	majority := data[0]
	cnt := 1
	for i := 1; i < len(data); {
		if majority != data[i] {
			cnt--
			i++

			if 0 == cnt && i < len(data) {
				majority = data[i]
				cnt++
				i++
			}
		} else {
			cnt++
			i++
		}
	}

	if 0 < cnt {
		return majority
	}

	return -1
}

//可能存在没有众数的情况
func MoreThanHalfNumber2(data []int) int {
	vote, x := 0, 0
	for _, v := range data {
		if 0 == vote {
			x = v
		}

		if v == x {
			vote++
		} else {
			vote--
		}
	}

	//verify
	cnt := 0
	for _, v := range data {
		if x == v {
			cnt++
		}
	}

	if cnt > len(data)/2 {
		return x
	}

	return -1
}
