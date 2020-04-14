package hashmap

import (
	"algorithmDataStructure/list/doublylinkedlist"
)

type Map struct {
	list *doublylinkedlist.List
	cap  int
	kvs  []*doublylinkedlist.List
	len  int
}

type KV struct {
	Key   int
	Value int
}

func New(cap int) *Map {
	return &Map{
		list: doublylinkedlist.New(),
		cap:  cap,
		kvs:  make([]*doublylinkedlist.List, cap),
	}
}

func (m *Map) Put(key int, value int) {

	list1 := m.kvs[m.hash(key)]
	if list1 == nil {
		e := &KV{Key: key, Value: value}
		m.kvs[m.hash(key)] = doublylinkedlist.New(e)
		m.list.PushBack(e)
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
		e := &KV{Key: key, Value: value}
		list1.PushBack(e)
		m.len++

		m.list.PushBack(e)
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
			m.list.Remove(e)
			m.len--
			return true
		}
		e = e.Next()
	}
	return false
}

//
func (m *Map) Len() int {
	return m.len
}

func (m *Map) hash(key int) int {
	return key % m.cap
}

func (m *Map) Iterator() *Iterator {
	return &Iterator{
		iterator: m.list.Iterator(),
	}
}

//
type Iterator struct {
	iterator *doublylinkedlist.Iterator
}

func (i *Iterator) Next() *KV {
	v, ok := i.iterator.Next()
	if !ok {
		return nil
	}
	return v.(*KV)
}
