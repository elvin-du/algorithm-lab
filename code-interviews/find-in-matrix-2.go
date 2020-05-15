package code_interviews

func FindInMatrix2(target int, matrix [][]int) bool {
	rl, rh := 0, len(matrix)-1
	cl, ch := 0, len(matrix[0])-1
	for ; rl <= rh && cl <= ch; {
		if target > matrix[rl][ch] {
			rl++
		} else if target < matrix[rl][ch] {
			ch--
		} else {
			return true
		}
	}

	return false
}
