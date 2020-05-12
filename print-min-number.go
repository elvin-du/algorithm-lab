//剑指offer 把数组排成最小的数
package algorithm_lab

import (
	"fmt"
	"sort"
	"strconv"
)

func PrintMinNumber(data []int) string {
	sort.Slice(data, func(i, j int) bool {
		a, err := strconv.ParseInt(fmt.Sprintf("%d%d", i, j), 10, 64)
		if nil != err {
			panic(err)
		}

		b, err := strconv.ParseInt(fmt.Sprintf("%d%d", j, i), 10, 64)
		if nil != err {
			panic(err)
		}

		return a > b
	})

	str := ""
	for _, v := range data {
		str = fmt.Sprintf("%s%d", str, v)
	}

	return str
}
