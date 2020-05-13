//剑指offer 数组中的逆序对
package algorithm_lab

var pairsCnt int = 0

func InversePairs(data []int) {
	mergeSort(data, 0, len(data)-1)
}

func mergeSort(data []int, l, h int) {
	if l < h {
		m := (h + l) / 2
		mergeSort(data, l, m)
		mergeSort(data, m+1, h)
		merge(data, l, m, h)
	}
}

func merge(data []int, l, m, h int) {
	tmp := make([]int, 0, h-l+1)
	left := l
	//因为除2得到的整数是舍去得来的，例如：3/2==1。所以这里已经是m+1作为后半部分的开始index，而不是m-1作为前半部分的结束index
	mid := m + 1
	for ; l <= m && mid <= h; {
		if data[l] < data[mid] {
			tmp = append(tmp, data[l])
			l++
		} else {
			tmp = append(tmp, data[mid])

			//计算前半部分的长度
			pairsCnt += m - l + 1
			mid++
		}
	}

	if l <= m {
		//a[left:right]right是开区间，不包含，所以要+1
		tmp = append(tmp, data[l:m+1]...)
	} else {
		tmp = append(tmp, data[mid:h+1]...)
	}

	for i, v := range tmp {
		data[left+i] = v
	}
}
