package code_interviews

func FindDuplicationNumberInArray(nums []int) int {
	for i, v := range nums {
		if v != nums[i] {
			nums[v], nums[i] = nums[i], nums[v]
		} else {
			return v
		}
	}

	return -1
}
