package hashmap

import (
	"algorithmDataStructure/list/doublylinkedlist"
)

//
type Map struct {
	cap int
	kvs []*doublylinkedlist.List
	len int
}

type KV struct {
	Key   int
	Value int
}

func New(cap int) *Map {
	return &Map{cap: cap, kvs: make([]*doublylinkedlist.List, cap)}
}

func (m *Map) Put(key int, value int) {

	list1 := m.kvs[m.hash(key)]
	if list1 == nil {
		m.kvs[m.hash(key)] = doublylinkedlist.New(&KV{Key: key, Value: value})
		m.len++
		return
	}

	e := list1.First

	have := false
	for e != nil {
		kv := e.Value.(*KV)

		if kv.Key == key {
			have = true
			kv.Value = value
		}
		e = e.Next()
	}

	if !have {
		list1.PushBack(&KV{Key: key, Value: value})
		m.len++
	}
}

func (m *Map) Get(key int) (int, bool) {
	list1 := m.kvs[m.hash(key)]
	if list1 == nil {
		return 0, false
	}
	e := list1.First
	for e != nil {
		kv := e.Value.(*KV)
		if kv.Key == key {
			return kv.Value, true
		}
		e = e.Next()
	}
	return 0, false
}

//
func (m *Map) Delete(key int) bool {
	list1 := m.kvs[m.hash(key)]
	if list1 == nil {
		return false
	}
	e := list1.First
	for e != nil {
		kv := e.Value.(*KV)
		if kv.Key == key {
			list1.Remove(e)
			m.len--
			return true
		}
		e = e.Next()
	}
	return false
}

func (m *Map) rehash() {
	m2 := New(m.cap * 2)
	for _, v := range m.kvs {
		if v == nil {
			continue
		}
		iter1 := v.Iterator()
		for {
			v, ok := iter1.Next()
			if !ok {
				break
			}
			m2.Put(v.(*KV).Key, v.(*KV).Value)
		}
	}
}

//
func (m *Map) Len() int {
	return m.len
}

func (m *Map) hash(key int) int {
	return key % m.cap
}
