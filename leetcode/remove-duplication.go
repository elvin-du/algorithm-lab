package leetcode

func DeleteDuplication(nums []int) int {
	j := 0

	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
