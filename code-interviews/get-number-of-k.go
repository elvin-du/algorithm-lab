package code_interviews

func GetNumberOfK(data []int, k int) int {
	l, h := 0, len(data)-1
	m := (l + h) / 2
	for ; l <= h; {
		m = (l + h) / 2
		if k > data[m] {
			l = m
		} else if k < data[m] {
			h = m
		} else {
			return numberOfK(data, m, k)
		}
	}

	return -1
}

func numberOfK(data []int, i, k int) int {
	cnt := 1
	for index := i+1; i <= len(data)-1; index++ {
		if k != data[index] {
			break
		}

		cnt++
	}

	for index := i-1; index >= 0; index-- {
		if k != data[index] {
			break
		}

		cnt++
	}

	return cnt
}
