//剑指offer 正则表达式匹配
package algorithm_lab

func RegexpMatch(i, j int, str, pattern string) bool {
	if i >= len(str) && j >= len(pattern) {
		return true
	} else if i >= len(str) || j >= len(pattern) {
		return false
	}

	if str[i] == pattern[j] || pattern[j] == '.' {
		if j+1 == len(pattern) && i+1 == len(str) {
			return true
		}

		if j+1 < len(pattern) && pattern[j+1] != '*' {
			return RegexpMatch(i+1, j+1, str, pattern)
		}

		return RegexpMatch(i+1, j, str, pattern) || RegexpMatch(i, j+2, str, pattern)
	} else if j+1 < len(pattern) && pattern[j+1] == '*' {
		return RegexpMatch(i, j+2, str, pattern)
	}

	//当前字符不匹配，且下一个模式字符不等于*，直接返回不匹配
	return false
}
