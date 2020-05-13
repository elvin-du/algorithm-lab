//剑指offer 反转单词顺序列
package algorithm_lab

func ReverseSentence(str string) string {
	data := []rune(str)
	l, h := 0, len(data)-1
	inverse(data, l, h)

	for left := l; l <= h; {
		if ' ' == data[l] {
			inverse(data, left, l-1)
			l++
			left = l
		} else if l+1 > h {
			inverse(data, left, l)
			l++
			left = l
		} else {
			l++
		}
	}

	return string(data)
}
