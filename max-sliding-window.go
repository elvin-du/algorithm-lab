//剑指offer 滑动窗口的最大值数组
package algorithm_lab

func MaxSlidingWindow(data []int, k int) []int {
	res := make([]int, 0, len(data)-k+1)
	dq := NewDequeue()

	for i := 0; i < len(data); i++ {
		for {
			if !dq.IsEmpty() {
				if data[dq.PeekBack().(int)] <= data[i] {
					dq.PollBack()
					continue
				}
			}

			break
		}

		dq.PushBack(i)

		if dq.PeekFront().(int) <= i-k {
			dq.PollFront()
		}

		if i >= k-1 {
			res = append(res, data[dq.PeekFront().(int)])
		}
	}

	return res
}
