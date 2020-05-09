//剑指offer 二进制1的个数
package algorithm_lab

func NumberOf1(n int) int {
	count := 0
	for ; n != 0; {
		count++
		n = n & (n - 1)
	}

	return count
}
