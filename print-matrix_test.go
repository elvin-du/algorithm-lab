package algorithm_lab

import "testing"

func TestPrintMatrix(t *testing.T) {
	data := [][]int{}
	row1 := []int{1, 2, 3, 4, 5}
	data = append(data, row1)
	row2 := []int{7, 8, 9, 10, 11}
	data = append(data, row2)

	row3 := []int{22, 33, 44, 55, 66}
	data = append(data, row3)

	PrintMatrix(data)
}
