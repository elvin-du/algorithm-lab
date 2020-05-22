package algorithm_lab

import (
	"algorithm-lab/common"
	"container/list"
)

type LRUCache struct {
	list *list.List
	cap  int
}

type LRUCacheNode struct {
	Key   int
	Value int
}

func NewLRUCacheNode(k, v int) *LRUCacheNode {
	return &LRUCacheNode{k, v}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{list: list.New(),  cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	v := this.get(key)
	if !common.IsNil(v) {
		val := v.Value.(*LRUCacheNode)
		this.list.MoveToBack(v)
		return val.Value
	}

	return -1
}
func (this *LRUCache) get(key int) *list.Element {
	head := this.list.Back()
	for ; !common.IsNil(head); {
		v := head.Value.(*LRUCacheNode)
		if v.Key == key {
			return head
		}

		head = head.Prev()
	}

	return nil
}

func (this *LRUCache) Put(key int, value int) {
	val := this.get(key)
	if common.IsNil(val) {
		if this.list.Len() +1 > this.cap {
			this.list.Remove(this.list.Front())
		}

		this.list.PushBack(NewLRUCacheNode(key, value))
	} else {
		val.Value = value
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor2(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
