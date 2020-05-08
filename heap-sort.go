package algorithm_lab

func HeapSort(data []int) {
	for i := len(data); i > 0; i-- {
		BuildMaxHeap(data, i)
		data[0], data[i-1] = data[i-1], data[0]
	}
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
