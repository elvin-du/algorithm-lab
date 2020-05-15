//剑指offer 构建乘积数组
package code_interviews

func Multiply(a []int) []int {
	n := len(a)
	b := make([]int, n)

	for i, product := 0, 1; i < n; i++ {
		b[i] = product
		product *= a[i]
	}

	for i, product := n-1, 1; i >= 0; i-- {
		b[i] *= product
		product *= a[i]
	}

	return b
}
