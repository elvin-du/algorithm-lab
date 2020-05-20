package code_interviews

func FindGreatestSumOfSubArray(data []int) int {
	max := data[0]
	sum := data[0]
	for i := 1; i < len(data); i++ {
		sum += data[i]
		if sum > max {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}
	return max
}
