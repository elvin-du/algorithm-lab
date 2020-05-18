package code_interviews

func GetMedian(data []int) int {
	panic("not impemented")
}

func BuildMaxHeap(data []int) {
	for i := len(data) / 2; i > 0; i-- {
		Heapify(data, i, len(data), maxHeapHelper)
	}
}

func InsertMaxHeap(data []int, num int) {
	data = append(data, num)
	BuildMaxHeap(data)
}

func HeapPoll(data []int) int {
	ret := data[0]
	data[0] = data[len(data)-1]
	BuildMaxHeap(data[:len(data)-1])
	return ret
}

func Heapify(data []int, k, size int, less func(int, int) bool) {
	index := k
	left := k*2 + 1
	right := k*2 + 2

	if left < size {
		if less(data[left], data[k]) {
			index = k
		}
	}

	if right < size {
		if less(data[right], data[k]) {
			index = k
		}
	}

	if index != k {
		data[index], data[k] = data[k], data[index]
		Heapify(data, index, size, less)
	}
}
func minHeapHelper(a, b int) bool {
	return a < b
}

func maxHeapHelper(a, b int) bool {
	return a > b
}
