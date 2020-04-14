package map1

import "fmt"

//开放定址法 - 线性探测法
//不带删除
type Map struct {
	kvs []*KV
	cap int
	len int
}

type KV struct {
	Key   int
	Value int
}

func New(cap int) *Map {
	return &Map{cap: cap, kvs: make([]*KV, cap)}
}

var (
	MapFull error = fmt.Errorf(`map full `)
)

func (m *Map) Put(key int, value int) error {
	if m.len == m.cap {
		return MapFull
	}

	x := -1
	for {
		x++
		if x >= m.cap {
			return MapFull
		}
		hashVal := m.hash(key + x)
		kv := m.kvs[hashVal]
		if kv == nil {
			m.kvs[hashVal] = &KV{Key: key, Value: value}
			m.len++
			return nil
		}
		if kv.Key == key {
			m.kvs[m.hash(key)].Value = value
			return nil
		}
	}
}

func (m *Map) Get(key int) (int, bool) {
	x := -1
	for {
		x++
		if x >= m.cap {
			return 0, false
		}
		hashVal := m.hash(key + x)
		kv := m.kvs[hashVal]
		if kv == nil {
			return 0, false
		}
		if kv.Key != key {
			continue
		}
		return kv.Value, true
	}
}

//
func (m *Map) Len() int {
	return m.len
}

func (m *Map) hash(key int) int {
	return key % m.cap
}
