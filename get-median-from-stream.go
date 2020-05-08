//剑指offer 从数据流中得到中间值
package algorithm_lab

var (
	maxHeap []int
	minHeap []int
)
//确保小根堆里面的数 > 大根堆里面的数
func InsertStreamData(number int) {
	if len(maxHeap) == len(minHeap) {
		//先从大顶堆过滤一遍,用来确保小顶堆里面的数是比大顶堆里面的数大
		maxHeap = InsertMaxHeap(maxHeap, number)
		max := MaxHeapPoll(&maxHeap)
		minHeap = InsertMinHeap(minHeap, max)
	} else {
		//先从小顶堆过滤一遍,用来确保大顶堆里面的数是比小顶堆里面的数小
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
