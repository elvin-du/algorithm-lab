//剑指offer 正则表达式匹配
package code_interviews

func RegexMatch(s, p int, str, pattern string) bool {
	if s > len(str)-1 && p > len(pattern)-1 {
		return true
	} else if s >= len(str) || p >= len(pattern) {
		return false
	}

	if str[s] == pattern[p] || '.' == pattern[p] {
		if len(pattern) == p+1 && len(str) == s+1 {
			return true
		}

		if len(pattern) > p+1 && '*' != pattern[p+1] {
			return RegexMatch(s+1, p+1, str, pattern)
		}

		return RegexMatch(s+1, p, str, pattern) || RegexMatch(s, p+2, str, pattern)
	} else if len(pattern) > p+1 && pattern[p+1] == '*' {
		return RegexMatch(s, p+2, str, pattern)
	}

	return false
}
