package lru2

import (
	"algorithmDataStructure/list/doublylinkedlist"
)

type LRUCache struct {
	m    map[int]*doublylinkedlist.Element
	list *doublylinkedlist.List
	cap  int
}

type KV struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:    make(map[int]*doublylinkedlist.Element, capacity),
		list: doublylinkedlist.New(),
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.m[key]
	if !ok {
		return -1
	}
	this.list.MoveToFirst(value)
	return value.Value.(KV).Value
}

func (this *LRUCache) Put(key int, value int) {
	element, ok := this.m[key]
	if ok {
		element.Value = KV{Key: key, Value: value}
		this.list.MoveToFirst(element)
	} else {
		e := this.list.PushFront(KV{Key: key, Value: value})
		this.m[key] = e
		l := len(this.m)
		if l > this.cap {
			back := this.list.Last
			delete(this.m, back.Value.(KV).Key)
			this.list.Remove(back)
		}
	}
}
