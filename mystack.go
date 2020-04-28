//leetcode 225
package algorithm_lab

type MyStack struct {
	queue1 *Queue
	queue2 *Queue //辅助
}

/** Initialize your data structure here. */
func NewMyStack() MyStack {
	return MyStack{NewQueue(), NewQueue()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	for ; !this.queue1.IsEmpty(); {
		this.queue2.Push(this.queue1.Pop())
	}

	this.queue1.Push(x)
	for ; !this.queue2.IsEmpty(); {
		this.queue1.Push(this.queue2.Pop())
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.queue1.Pop().(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue1.Peek().(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue1.IsEmpty()
}

/**
* Your MyStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.Empty();
 */
