//剑指offer 把空格替换成%20
package algorithm_lab

func ReplaceSpace(str string) string {
	count := CountSpaceRecursion(0, str)
	replace := make([]byte, count*2+len(str))
	for i, j := len(str)-1, len(replace)-1; i >= 0; i-- {
		if ' ' != str[i] {
			replace[j] = str[i]
			j--
		} else {
			j -= 3
			copy(replace[j+1:], []byte{'%', '2', '0'})
		}
	}

	return string(replace)
}

func CountSpaceRecursion(index int, str string) int {
	if index >= len(str) {
		return 0
	}

	if str[index] == ' ' {
		return 1 + CountSpaceRecursion(index+1, str)
	}

	return CountSpaceRecursion(index+1, str)
}

func CountSpace(str string) int {
	count := 0
	for _, v := range str {
		if ' ' == v {
			count++
		}
	}

	return count
}
