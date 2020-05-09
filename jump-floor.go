//剑指offer 跳台阶
package algorithm_lab

func numWays(n int) int {
	if n == 0 {
		return 1
	}

	if n <= 2 {
		return n
	}

	return numWays(n-1) + numWays(n-2)
}

func numWays1(n int) int {
	if n <= 1 {
		return 1
	}

	f0, f1 := 1, 1
	for i := 2; i < n; i++ {
		tmp := f1
		f1 = f1 + f0
		f0 = tmp
	}

	return f1 + f0
}
