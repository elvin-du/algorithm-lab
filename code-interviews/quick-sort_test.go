package code_interviews

import "testing"

func TestQuickSort(t *testing.T) {
	data := []int{4,8,1,7,9,6}
	QuickSort(data)
	t.Log(data)
}