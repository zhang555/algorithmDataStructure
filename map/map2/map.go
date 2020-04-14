package map1

import "fmt"

//开放定址法 - 线性探测， 带删除
type Map struct {
	kvs []KV
	cap int
	len int
}

type KV struct {
	Key    int
	Value  int
	Status Status
}

type Status int8

const (
	Empty  Status = 0
	Exist  Status = 1
	Delete Status = 2
)

func New(cap int) *Map {
	return &Map{cap: cap, kvs: make([]KV, cap)}
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

		if kv.Status == Delete || kv.Status == Empty {
			m.kvs[hashVal].Status = Exist
			m.kvs[hashVal].Key = key
			m.kvs[hashVal].Value = value
			m.len++
			return nil
		}
		if kv.Status == Exist {
			if kv.Key != key {
				continue
			}
			m.kvs[hashVal].Value = value
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
		if kv.Status == Delete {
			continue
		}
		if kv.Status == Empty {
			return 0, false
		}
		if kv.Key != key {
			continue
		}
		return kv.Value, true
	}

	//pos := m.getPos(key)
	//if pos == -1 {
	//	return 0, false
	//}
	//return m.kvs[pos].Value, true

}

func (m *Map) Delete(key int) {
	x := -1
	for {
		x++
		if x >= m.cap {
			return
		}
		hashVal := m.hash(key + x)
		kv := m.kvs[hashVal]
		if kv.Status == Delete {
			continue
		}
		if kv.Status == Empty {
			return
		}
		if kv.Key != key {
			continue
		}
		m.kvs[hashVal].Status = Delete
		m.len--
		return
	}
}

//func (m *Map) getPos(key int) int {
//	x := -1
//	for {
//		x++
//		if x >= m.cap {
//			return -1
//		}
//		hashVal := m.hash(key + x)
//		kv := m.kvs[hashVal]
//		if kv.Status == Delete {
//			continue
//		}
//		if kv.Status == Empty {
//			return -1
//		}
//		if kv.Key != key {
//			continue
//		}
//		return hashVal
//	}
//}

//
func (m *Map) Len() int {
	return m.len
}

func (m *Map) hash(key int) int {
	return key % m.cap
}
