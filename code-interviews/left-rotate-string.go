package code_interviews

import "bytes"

func LeftRotateString(str string, k int) string {
	s1 := bytes.Runes([]byte(str))
	s := make([]rune,len(s1))
	copy(s,s1)
	left := s[:k]
	right := s[k:]
	rotate(left)
	rotate(right)
	rotate(s)
	return string(s)
}

func rotate(s []rune) {
	l, h := 0, len(s)-1
	for ; l <= h; {
		s[l], s[h] = s[h], s[l]
		l++
		h--
	}
}
