package algorithm_lab

func HeapSort(data []int) {
	for i := len(data); i > 0; i-- {
		BuildMaxHeap(data, i)
		data[0], data[i-1] = data[i-1], data[0]
	}
}

func MaxHeapPoll(data *[]int) int {
	i := len(*data)
	ret := (*data)[0]

	(*data)[0] = (*data)[i-1]
	BuildMaxHeap(*data, i-1)
	*data = (*data)[:i-1]

	return ret
}

func BuildMaxHeap(data []int, size int) {
	for l := size/2 - 1; l >= 0; l-- {
		Heapify(data, l, size)
	}
}

func InsertMaxHeap(data []int, newData int) []int {
	data = append(data, newData)
	BuildMaxHeap(data, len(data))
	return data
}

func Heapify(data []int, index, size int) {
	left := index*2 + 1
	right := index*2 + 2
	largestIndex := index

	if left < size {
		if data[left] > data[largestIndex] {
			largestIndex = left
		}
	}

	if right < size {
		if data[right] > data[largestIndex] {
			largestIndex = right
		}
	}

	if largestIndex != index {
		data[largestIndex], data[index] = data[index], data[largestIndex]
		Heapify(data, largestIndex, size)
	}
}

/***************************************
小根堆
*****************************************/
func MinHeapSort(data []int) {
	for i := len(data); i > 0; i-- {
		BuildMinHeap(data, i)
		data[0], data[i-1] = data[i-1], data[0]
	}
}

func BuildMinHeap(data []int, size int) {
	for l := size/2 - 1; l >= 0; l-- {
		MinHeapify(data, l, size)
	}
}

func InsertMinHeap(data []int, newData int) []int {
	data = append(data, newData)
	BuildMinHeap(data, len(data))
	return data
}

func MinHeapify(data []int, index, size int) {
	left := index*2 + 1
	right := index*2 + 2
	largestIndex := index

	if left < size {
		if data[left] < data[largestIndex] {
			largestIndex = left
		}
	}

	if right < size {
		if data[right] < data[largestIndex] {
			largestIndex = right
		}
	}

	if largestIndex != index {
		data[largestIndex], data[index] = data[index], data[largestIndex]
		MinHeapify(data, largestIndex, size)
	}
}

func MinHeapPoll(data *[]int) int {
	i := len(*data)
	ret := (*data)[0]

	(*data)[0] = (*data)[i-1]
	BuildMinHeap(*data, i-1)
	*data = (*data)[:i-1]

	return ret
}
