package algorithm_lab

type Queue struct {
	List *List
}

func NewQueue() *Queue {
	return &Queue{List: NewList()}
}

func (q *Queue) Push(e interface{}) {
	q.List.PushBack(e)
}

func (q *Queue) Pop() interface{} {
	return q.List.RemoveByIndex(0)
}

func (q *Queue) Peek() interface{} {
	return q.List.Get(0)
}

func (q *Queue) Size() int {
	return q.List.Size
}

func (q *Queue) IsEmpty() bool {
	return q.List.Size == 0
}
