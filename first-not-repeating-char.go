//剑指offer 第一个只出现一次的字符
package algorithm_lab

func FirstNotRepeatingChar(str string) int {
	words := [58]int{}
	for _, r := range str {
		words[r-65] += 1
	}

	for i, r := range str {
		if words[r-65] == 1 {
			return i
		}
	}

	return -1
}
