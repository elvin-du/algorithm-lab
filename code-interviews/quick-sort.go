package code_interviews

func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, l, h int) {
	if h <= l {
		return
	}

	p := data[l]
	left, right := l, h
	for ; l < h; {
		for ; l < h; {
			if data[h] < p {
				data[h], data[l] = data[l], data[h]
				l++
				break
			} else {
				h--
			}
		}

		for ; l < h; {
			if data[l] > p {
				data[h], data[l] = data[l], data[h]
				h--
				break
			} else {
				l++
			}
		}
	}

	quickSort(data, left, h)
	quickSort(data, h+1, right)
}
