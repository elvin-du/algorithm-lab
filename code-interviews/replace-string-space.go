//剑指offer 替换空格
package code_interviews

func ReplaceStringSpace(str string) string {
	spaceCnt := 0
	for _, v := range str {
		if ' ' == v {
			spaceCnt++
		}
	}
	newStrSize := spaceCnt*2 + len(str)
	newStr := make([]rune, newStrSize)
	newStrIndex := newStrSize - 1
	runeArr := []rune(str)

	for i := len(runeArr) - 1; i >= 0; i--{
		if str[i] == ' ' {
			newStr[newStrIndex] = '0'
			newStr[newStrIndex-1] = '2'
			newStr[newStrIndex-2] = '%'
			newStrIndex -= 3
		} else {
			newStr[newStrIndex] = runeArr[i]
			newStrIndex--
		}
	}

	return string(newStr)
}
