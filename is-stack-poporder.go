//剑指offer 栈的压入和弹出序列
package algorithm_lab

func IsStackPopOrder(pushSequence, popSequence []int) bool {
	stack := NewStack(100)
	i := 0
	for p := 0; p < len(popSequence); {
		if stack.Size() != 0 {
			if popSequence[p] == stack.Top().(int) {
				stack.Pop()
				p++
				continue
			}
		}

		if i >= len(pushSequence) {
			break
		}

		for ; i < len(pushSequence); i++ {
			if popSequence[p] != pushSequence[i] {
				stack.Push(pushSequence[i])
			} else {
				p++
				i++
				break
			}
		}
	}

	if i >= len(pushSequence) && stack.Size() != 0 {
		return false
	}

	return true
}
