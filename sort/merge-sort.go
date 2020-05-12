//归并排序
package sort

func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	//分
	mid := len(data) / 2
	left, right := data[:mid], data[mid:]

	return Merge(MergeSort(left), MergeSort(right))
}

//治
func Merge(a, b []int) []int {
	c := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for ; i < len(a) && j < len(b); {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	if i < len(a) {
		c = append(c, a[i:]...)
	}

	if j < len(b) {
		c = append(c, b[j:]...)
	}

	return c
}
