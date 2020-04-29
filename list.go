package algorithm_lab

import (
	"fmt"
	"reflect"
)

type List struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

func NewList() *List {
	return &List{}
}

func (l *List) String() string {
	str := ""
	node := l.Head
	for ; node != nil; node = node.Next {
		str += fmt.Sprintf("%v", node.Value)
	}
	return str
}

func (l *List) PushBack(e interface{}) {
	newNode := NewListNode(e)
	l.Size += 1
	if nil == l.Tail && nil == l.Head {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	l.Tail.Next = newNode
	newNode.Pre = l.Tail
	l.Tail = l.Tail.Next
}

func (l *List) PushFront(e interface{}) {
	newNode := NewListNode(e)
	l.Size += 1
	if nil == l.Tail && nil == l.Head {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	oldHead := l.Head
	l.Head = newNode
	newNode.Next = oldHead

	oldHead.Pre = newNode
}

//index不可以超过链表最大长度
func (l *List) Set(index int, e interface{}) {
	if index > l.Size {
		panic("index out of range")
	}

	if index == l.Size {
		l.PushBack(e)
		return
	}

	newNode := NewListNode(e)
	node := l.Head
	//target的前一个元素
	for i := 0; i < index-1; i++ {
		node = node.Next
	}

	target := node.Next

	node.Next = newNode
	newNode.Pre = node

	newNode.Next = target
	target.Pre = newNode
}

func (l *List) Get(index int) interface{} {
	if index >= l.Size {
		panic("index out of range")
	}

	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node.Value
}

//return -1 意味着没有找到
func (l *List) Find(e interface{}) int {
	i := 0
	for n := l.Head; nil != n; n = n.Next {
		if reflect.DeepEqual(n.Value, e) {
			return i
		}
		i += 1
	}

	return -1
}

func (l *List) RemoveByIndex(index int) interface{} {
	if index >= l.Size {
		return nil
	}

	//第一个元素
	if 0 == index && nil != l.Head {
		target := l.Head
		l.Head = l.Head.Next
		if nil != l.Head {
			l.Head.Pre = nil
		} else {
			l.Tail = nil //没有节点了
		}
		l.Size -= 1
		return target.Value
	}

	//最后一个元素, 而且链表中的中的元素最少也是 >= 2，这一点是通过前面一个if语句来保证
	if index == l.Size-1 {
		target := l.Tail
		l.Tail = target.Pre
		target.Pre = nil
		l.Tail.Next = nil
		return target.Value
	}

	i := -1
	for n := l.Head; n != nil; n, i = n.Next, i+1 {
		if i == index-1 {

			target := n.Next
			n.Next = target.Next
			target.Next.Pre = n
			l.Size -= 1

			return target.Value
		}
	}

	return nil
}

type ListNode struct {
	Next  *ListNode
	Pre   *ListNode
	Value interface{}
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{Value: v}
}
