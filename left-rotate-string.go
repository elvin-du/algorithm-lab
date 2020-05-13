//剑指offer 左移字符串
package algorithm_lab

import "bytes"

func LeftRotateString(str string, k int) string {
	arr := bytes.Runes([]byte(str))
	inverse(arr, 0, k-1)
	inverse(arr, k, len(arr)-1)
	inverse(arr, 0, len(arr)-1)
	return string(arr)
}

func inverse(arr []rune, l, h int) {
	for ; l < h; {
		arr[l], arr[h] = arr[h], arr[l]
		l++
		h--
	}
}
