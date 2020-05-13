//剑指offer 和为S的两个数字
package algorithm_lab

import "math"

//data：递增数组
func FindNumbersWithSum(data []int, s int) [2]int {
	l, h := 0, len(data)-1
	minMul := math.MaxInt64
	minMulArr := [2]int{}
	for ; l <= h; {
		sum := data[l] + data[h]
		if sum == s {
			mul := data[l] * data[h]
			if mul < minMul {
				minMul = mul
				minMulArr[0], minMulArr[1] = data[l], data[h]
			}

			l++
			h--
		} else if sum < s {
			l++
		} else {
			h--
		}
	}

	return minMulArr
}
