//leetcode 155
package algorithm_lab

import "algorithm-lab/common"

type MinStack struct {
	min   int
	stack *common.Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	s := MinStack{stack: common.NewStack(100)}
	return s
}

func (this *MinStack) Push(x int) {
	if this.stack.Size() == 0 {
		this.min = x
		this.stack.Push(x)
		return
	}

	//这里是关键；用于存储上一次的最小值。
	//如果相等也需要存储，因为pop的时候，需要弹出
	if x <= this.min {
		this.stack.Push(this.min)
		this.stack.Push(x)
		this.min = x
		return
	}

	this.stack.Push(x)
}

func (this *MinStack) Pop() {
	x := this.stack.Pop().(int)
	if x == this.min {
		//已经到底了
		if this.stack.Size() == 0 {
			return
		}

		x2 := this.stack.Pop().(int)
		this.min = x2
	}
}

func (this *MinStack) Top() int {
	return this.stack.Top().(int)
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
