//剑指offer 数字在排序数组中的次数
package algorithm_lab

func GetNumberOfK(data []int, k int) int {
	return GetLastKIndex(data, k) - GetFirstKIndex(data, 0, len(data)-1, k) + 1
}

func GetFirstKIndex(data []int, l, h, k int) int {
	if l > h {
		return -1
	}

	m := l + (h-l)/2
	if data[m] > k {
		return GetFirstKIndex(data, l, m-1, k)
	} else if data[m] < k {
		return GetFirstKIndex(data, m+1, h, k)
	} else if m-1 >= 0 && data[m-1] == k {
		return GetFirstKIndex(data, l, m-1, k)
	}

	return m
}

func GetLastKIndex(data []int, k int) int {
	l, h := 0, len(data)-1
	for ; l <= h; {
		m := l + (h-l)/2
		if data[m] > k {
			h = m - 1
		} else if data[m] < k {
			l = m + 1
		} else if m+1 <= h && data[m+1] == k {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}
