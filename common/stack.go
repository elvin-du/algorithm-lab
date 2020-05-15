package common

type Stack struct {
	Data     []interface{}
	TopIndex int //指向最新元素的索引
	MaxDepth int
}

func NewStack(maxDepth int) *Stack {
	return &Stack{Data: make([]interface{}, maxDepth), TopIndex: -1, MaxDepth: maxDepth}
}

func (s *Stack) Pop() interface{} {
	e := s.Top()
	s.TopIndex -= 1
	return e
}

func (s *Stack) Push(e interface{}) {
	s.TopIndex += 1
	s.Data[s.TopIndex] = e
}

func (s *Stack) Top() interface{} {
	if s.TopIndex < 0 {
		panic("Stack is empty")
	}
	return s.Data[s.TopIndex]
}

func (s *Stack) Size() int {
	return s.TopIndex + 1
}
