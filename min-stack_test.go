package algorithm_lab

//
//func TestMinStack(t *testing.T) {
//	minStack := Constructor()
//	minStack.Push(8)
//	minStack.Push(9)
//	minStack.Push(5)
//	minStack.Push(7)
//	minStack.Push(1)
//
//	assert.Equal(t, 1, minStack.GetMin())
//	minStack.Pop()
//	assert.Equal(t, 5, minStack.GetMin())
//}
//
//func TestMinStack2(t *testing.T) {
//	minStack := Constructor()
//	minStack.Push(0)
//	minStack.Push(1)
//	minStack.Push(0)
//
//	assert.Equal(t, 0, minStack.GetMin())
//	minStack.Pop()
//	assert.Equal(t, 0, minStack.GetMin())
//}
//
//func TestMinStack3(t *testing.T) {
//	minStack := Constructor()
//	minStack.Push(2147483646)
//	minStack.Push(2147483646)
//	minStack.Push(2147483647)
//
//	assert.Equal(t, 2147483647, minStack.Top())
//	minStack.Pop()
//	assert.Equal(t, 2147483646, minStack.GetMin())
//	minStack.Pop()
//
//	assert.Equal(t, 2147483646, minStack.GetMin())
//
//	minStack.Pop()
//
//	minStack.Push(2147483647)
//	assert.Equal(t, 2147483647, minStack.Top())
//
//	assert.Equal(t, 2147483647, minStack.GetMin())
//	minStack.Push(-2147483648)
//	assert.Equal(t, -2147483648, minStack.Top())
//	assert.Equal(t, -2147483648, minStack.GetMin())
//
//	minStack.Pop()
//	assert.Equal(t, 2147483647, minStack.GetMin())
//}
