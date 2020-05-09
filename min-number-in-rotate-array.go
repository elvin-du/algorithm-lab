//剑指offer 旋转数组中的最小值
package algorithm_lab

func MinNumberInRotateArray(data []int) int {
	left, right := 0, len(data)-1
	mid := (right - left) / 2

	for {
		if data[mid] > data[left] {
			left = mid
			mid = (right + left) / 2
		} else if data[right] > data[mid] {
			right = mid
			mid = (right + left) / 2
		}

		if left+1 == right {
			return data[right]
		}
	}
}
