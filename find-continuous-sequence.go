//剑指offer 和为S的连续序列
package algorithm_lab

func FindContinuousSequence(n int) [][]int {
	//两个起点，相当于动态窗口的两边，根据其窗口内的值的和来确定窗口的位置和大小
	l, h := 1, 2
	res := [][]int{}
	for l < h && h < n/2 {
		sum := Sum(l, h)
		//相等，那么就将窗口范围的所有数添加进结果集
		if sum == n {
			arr := []int{}
			for i := l; i <= h; i++ {
				arr = append(arr, i)
			}
			res = append(res, arr)
			l++
		} else if sum < n {
			//如果当前窗口内的值之和小于sum，那么右边窗口右移一下
			h++
		} else {
			//如果当前窗口内的值之和大于sum，那么左边窗口右移一下
			l++
		}
	}

	return res
}

func Sum(l, h int) int {
	if l > h {
		return 0
	}

	return l + Sum(l+1, h)
}
