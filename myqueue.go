package algorithm_lab

import "algorithm-lab/common"

type MyQueue struct {
	stack1 *common.Stack
	stack2 *common.Stack
}

/** Initialize your data structure here. */
func NewMyQueue() MyQueue {
	return MyQueue{common.NewStack(1000), common.NewStack(1000)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	size1 := this.stack1.Size()
	for i := 0; i < size1; i++ {
		this.stack2.Push(this.stack1.Pop())
	}
	this.stack1.Push(x)

	size2 := this.stack2.Size()
	for i := 0; i < size2; i++ {
		this.stack1.Push(this.stack2.Pop())
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return this.stack1.Pop().(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack1.Top().(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack1.Size() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
