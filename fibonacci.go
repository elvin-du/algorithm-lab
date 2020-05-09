//剑指offer fibonacci
package algorithm_lab

func Fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci2(n int) int {
	if n <= 1 {
		return 1
	}

	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func Fibonacci3(n int) int {
	if n <= 1 {
		return 1
	}

	f0, f1 := 1, 1

	for i := 2; i < n; i++ {
		tmp := f1
		f1 = f0 + f1
		f0 = tmp
	}

	return f0 + f1
}
