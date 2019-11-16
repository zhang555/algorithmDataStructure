package lru

import (
	"container/list"
)

type LRUCache struct {
	m    map[int]*list.Element
	list *list.List
	cap  int
}

type KV struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:    make(map[int]*list.Element, capacity),
		list: list.New(),
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.m[key]
	if !ok {
		return -1
	}
	this.list.MoveToFront(value)
	return value.Value.(KV).Value
}

func (this *LRUCache) Put(key int, value int) {
	element, ok := this.m[key]
	if ok {
		element.Value = KV{Key: key, Value: value}
		this.list.MoveToFront(element)
	} else {
		e := this.list.PushFront(KV{Key: key, Value: value})
		this.m[key] = e
		l := len(this.m)
		if l > this.cap {
			back := this.list.Back()
			delete(this.m, back.Value.(KV).Key)
			this.list.Remove(back)
		}
	}
}
