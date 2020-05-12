//剑指offer 最小的K个数
package algorithm_lab

func GetKLeastNumbers(data []int, k int) []int {
	FindKthSmallest(data, k)
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, data[i])
	}
	return res
}

func FindKthSmallest(data []int, k int) {
	l, h := 0, len(data)-1
	for ; l < h; {
		index := GetPartitionIndex(data, l, h)
		if index == k {
			break
		}

		if index > k {
			l = index + 1
		} else {
			h = index - 1
		}
	}
}

func GetPartitionIndex(data []int, l, h int) int {
	pivot := data[l]
	for ; l < h; {
		if data[h] > pivot {
			h--
			continue
		}

		if data[l] < pivot {
			l++
			continue
		}

		data[l], data[h] = data[h], data[l]
	}

	return l
}
