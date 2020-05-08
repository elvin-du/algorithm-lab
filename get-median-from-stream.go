//剑指offer 从数据流中得到中间值
package algorithm_lab

var (
	maxHeap []int
	minHeap []int
)

func InsertStreamData(number int) {
	if len(maxHeap) == len(minHeap) {
		maxHeap = InsertMaxHeap(maxHeap, number)
		max := MaxHeapPoll(&maxHeap)
		minHeap = InsertMinHeap(minHeap, max)
	} else {
		minHeap = InsertMinHeap(minHeap, number)
		min := MinHeapPoll(&minHeap)
		maxHeap = InsertMaxHeap(maxHeap, min)
	}
}

func GetMedian() int {
	if len(minHeap) != len(maxHeap) {
		return minHeap[0]
	}

	return (minHeap[0] + maxHeap[0]) / 2
}
