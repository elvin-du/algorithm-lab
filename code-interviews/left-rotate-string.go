package code_interviews

import "bytes"

func LeftRotateString(str string, k int) string {
	s := bytes.Runes([]byte(str))
	left := s[:k]
	right := s[k:]
	rotate(left)
	rotate(right)
	rotate(s)
	return str
}

func rotate(s []rune) {
	l, h := 0, len(s)-1
	for ; l <= h; {
		s[l], s[h] = s[h], s[l]
		l++
		h--
	}
}
