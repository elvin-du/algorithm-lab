//剑指offer 整数中1出现的次数
package algorithm_lab

func NumberOf1Between1AndN(n int) int {
	cnt := 0
	for i := 1; i <= n; i++ {
		cnt += count1(i)
	}

	return cnt
}

func count1(n int) int {
	cnt := 0
	for ; n != 0; {
		if n%10 == 1 {
			cnt++
		}

		n = n / 10
	}

	return cnt
}
