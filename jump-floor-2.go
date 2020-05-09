//剑指offer 变态跳台阶
package algorithm_lab

func JumpFloor2(n int) int {
	f0 := 1

	for i := 1; i <= n; i++ {
		f0 *= 2
	}

	return f0
}
