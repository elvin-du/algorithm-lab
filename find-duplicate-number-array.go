//剑指offer 从数组中找出第一个重复的数字
package algorithm_lab

func FindDuplicateNumberInArray(arr []int) int {
	for i := 0; i < len(arr); {
		if i == arr[i] {
			i++
			continue
		} else {
			if Swap(i, arr) {
				return arr[i]
			}
		}
	}

	return -1
}

func Swap(index int, arr []int) bool {
	if arr[index] == arr[arr[index]] {
		return true
	}
	arr[index], arr[arr[index]] = arr[arr[index]], arr[index]

	return false
}
