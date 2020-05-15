//剑指offer 滑动窗口的最大值数组
package algorithm_lab

import "algorithm-lab/common"

func MaxSlidingWindow(data []int, k int) []int {
	res := make([]int, 0, len(data)-k+1)
	dq := common.NewDequeue()

	for i := 0; i < len(data); i++ {
		//把小值弹出双端单调递减队列
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

		//是否已经超出滑动窗口
		if dq.PeekFront().(int) <= i-k {
			dq.PollFront()
		}

		if i >= k-1 {
			res = append(res, data[dq.PeekFront().(int)])
		}
	}

	return res
}
