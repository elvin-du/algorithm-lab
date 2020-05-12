//快速排序
package sort

func QuickSort(data []int) {
	Partition(data, 0, len(data)-1)
	//p := Partition(data, 0, len(data)-1)
	//fmt.Println(p)
}

func Partition(data []int, l, h int) int {
	if l >= h {
		return h
	}

	low, high := l, h

	pivot := data[l]
	for ; l < h; {
		if data[h] > pivot {
			h--
			continue
		}

		if data[l] < pivot {
			l++
			continue
		}

		//swap
		data[h], data[l] = data[l], data[h]
	}

	Partition(data, low, l-1)
	Partition(data, l+1, high)

	return l
}
