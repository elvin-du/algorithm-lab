//剑指offer 连续子数组的最大和
package algorithm_lab

func FindGreatestSumOfSubArray(data []int) int {
	max := data[0]
	sum := 0
	for _, v := range data {
		sum += v
		if sum > max {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return max
}
