//剑指offer 字符流中第一个不重复的字符
package algorithm_lab

import (
	"algorithm-lab/common"
	"io"
)

func FirstAppearingOnce(r io.Reader) byte {
	m := [256]byte{}
	b := make([]byte, 1)
	queue := common.NewQueue()

	for ; ; {
		_, err := r.Read(b)
		if nil != err {
			//log.Println("read err:", err)
			break
		}

		m[b[0]] += 1
		if m[b[0]] == 1 {
			queue.Push(b[0])
		}
	}

	for ; !queue.IsEmpty(); {
		v := queue.Pop().(byte)
		if m[v] == 1 {
			return v
		}
	}

	return 0
}
