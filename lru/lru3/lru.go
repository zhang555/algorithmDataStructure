package lru3

import "container/list"

type LRU struct {
	m   map[int]*list.Element
	l   *list.List
	cap int
}

type KV struct {
	Key   int
	Value int
}

func Constructor(cap int) *LRU {
	return &LRU{
		m:   map[int]*list.Element{},
		l:   list.New(),
		cap: cap,
	}
}

func (lru *LRU) Get(key int) int {
	v, ok := lru.m[key]
	if !ok {
		return -1
	}
	lru.l.MoveToFront(v)
	return v.Value.(*KV).Value
}

func (lru *LRU) Put(key int, value int) {
	v, ok := lru.m[key]
	if !ok {
		if len(lru.m) == lru.cap {
			element := lru.l.Back()
			lru.l.Remove(lru.l.Back())

			delete(lru.m, element.Value.(*KV).Key)

			element = lru.l.PushFront(&KV{Key: key, Value: value})
			lru.m[key] = element
			return
		}

		element := lru.l.PushFront(&KV{Key: key, Value: value})
		lru.m[key] = element
		return
	}

	lru.l.MoveToFront(v)
	v.Value.(*KV).Value = value
	return
}
