//剑指offer string转成整数
package algorithm_lab

func StrToInt(str string) int {
	num := 0
	minus := false
	for i, v := range str {
		if i != 0 && !isNumber(v) {
			return 0
		}

		if i == 0 && v == '-' {
			minus = true
		} else if isNumber(v) {
			num *= 10
			num += int(v - '0')
		}
	}

	if minus{
		return -num
	}

	return num
}

func isNumber(v rune) bool {
	return v >= '0' && v <= '9'
}
