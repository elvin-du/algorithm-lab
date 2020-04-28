//剑指offer - 二维数组查找
package algorithm_lab

func ExistInMatrix(target int, matrix [][]int) bool {
	rows := len(matrix)
	columns := len(matrix[0])

	r := 0           //从第一行开始向下搜索
	c := columns - 1 //从最后一列向前开始搜索
	for ; r < rows; r++ {
		for ; c >= 0; c-- {
			if matrix[r][c] < target {
				break
			} else if matrix[r][c] == target {
				return true
			}
		}
	}

	return false
}
