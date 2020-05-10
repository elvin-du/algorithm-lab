//剑指offer 调整数组顺序使奇数位于偶数前面
package algorithm_lab

func ReorderArray(nums []int) {
	left, right := 0, len(nums)-1

	for ; left < right; {
		if nums[left]%2 == 1 {
			left++
		}
		if nums[right]%2 == 0 {
			right--
		}

		if nums[left]%2 == 0 && nums[right]%2 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
}
