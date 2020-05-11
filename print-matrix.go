//剑指offer 顺时针打印矩阵
package algorithm_lab

func PrintMatrix(data [][]int) {
	r1, r2 := 0, len(data)-1
	c1, c2 := 0, len(data[0])-1
	list := NewList()

	for ; r1 <= r2 && c1 <= c2; {
		for i := c1; i <= c2; i++ {
			list.PushBack(data[r1][i])
		}

		for i := r1 + 1; i <= r2; i++ {
			list.PushBack(data[i][c2])
		}

		if r1 != r2 {
			for i := c2 - 1; i >= c1; i-- {
				list.PushBack(data[r2][i])
			}
		}

		if c1 != c2 {
			for i := r2 - 1; i > r1; i-- {
				list.PushBack(data[i][c1])
			}
		}

		r1++
		r2--
		c1++
		c2--
	}

	list.Print()
}
