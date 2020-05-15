//leetcode 20
package algorithm_lab

import "algorithm-lab/common"

func isValid(s string) bool {
	stack := common.NewStack(len(s))
	for _, right := range s {
		if 0 < stack.Size() {
			left := stack.Top()
			if IsMatch(left.(int32), right) {
				stack.Pop()
				continue
			}
		}
		stack.Push(right)
	}

	return stack.Size() == 0
}

func IsMatch(left, right int32) bool {
	switch left {
	case '(':
		if right == ')' {
			return true
		}
	case '[':
		if right == ']' {
			return true
		}
	case '{':
		if right == '}' {
			return true
		}
	}

	return false
}
