//剑指offer 字符流中第一个不重复的字符
package code_interviews

import (
	"algorithm-lab/common"
	"io"
)

func FirstAppearingOnce(r io.Reader) string {
	b := make([]byte, 1)
	m := map[string]int{}
	q := common.NewQueue()
	for {
		_, err := r.Read(b)
		if nil != err {
			break
		}

		m[string(b[0])] += 1
		if m[string(b[0])] == 1 {
			q.Push(string(b[0]))
		}
	}
	for {
		if !q.IsEmpty() {
			v := q.Pop().(string)
			if 1 == m[v] {
				return v
			}
		}
	}

	return ""
}
