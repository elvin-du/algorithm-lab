package common

type Dequeue struct {
	list *List
}

func NewDequeue() *Dequeue {
	return &Dequeue{NewList()}
}

func (dq *Dequeue) PushBack(e interface{}) {
	dq.list.PushBack(e)
}

func (dq *Dequeue) PushFront(e interface{}) {
	dq.list.PushFront(e)
}

func (dq *Dequeue) PollFront() interface{} {
	return dq.list.RemoveByIndex(0)
}

func (dq *Dequeue) PollBack() interface{} {
	return dq.list.RemoveByIndex(dq.list.Size - 1)
}

func (dq *Dequeue) PeekBack() interface{} {
	return dq.list.Tail.Value
}

func (dq *Dequeue) PeekFront() interface{} {
	return dq.list.Head.Value
}

func (dq *Dequeue) IsEmpty() bool {
	return dq.list.Size == 0
}

func (dq *Dequeue) Size() int {
	return dq.list.Size
}
