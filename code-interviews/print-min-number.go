package code_interviews

import (
	"fmt"
	"sort"
	"strconv"
)

func PrintMinNumber(data []int) string {
	sort.Slice(data, func(i, j int) bool {
		a, _ := strconv.ParseUint(fmt.Sprintf("%d%d", i, j), 10, 64)
		b, _ := strconv.ParseUint(fmt.Sprintf("%d%d", j, i), 10, 64)
		return a > b
	})

	str := ""
	for _, v := range data {
		str += fmt.Sprintf("%d", v)
	}
	return str
}
