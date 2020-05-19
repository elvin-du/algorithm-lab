package code_interviews

import "algorithm-lab/common"

func MaxInWindows(data []int, k int) []int {
	q := common.NewDequeue()
	max := []int{}
	for i := 0; i < len(data); i++ {
		for {
			if q.IsEmpty() {
				q.PushBack(i)
				break
			}

			index := q.PeekBack().(int)
			if data[index] > data[i] {
				q.PushBack(i)
				break
			}

			q.PollBack()
		}

		if !q.IsEmpty() {
			index := q.PeekFront().(int)
			if index <= i-k {
				q.PollFront()
			}
		}

		if i >= k-1 {
			max = append(max, data[q.PeekFront().(int)])
		}
	}

	return max
}
