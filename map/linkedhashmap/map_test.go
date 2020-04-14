package hashmap

import (
	"math/rand"
	"testing"
	"time"
)

func TestMap000(t *testing.T) {
	m := New(1000)

	m.Put(1, 1)

	v, ok := m.Get(1)
	if !ok {
		t.Error()
	}
	if v != 1 {
		t.Error()
	}

}

func TestMap001(t *testing.T) {
	m := New(1000)

	for i := 0; i < 100; i++ {
		m.Put(i, i)
		v, ok := m.Get(i)
		if !ok {
			t.Error()
		}
		if v != i {
			t.Error()
		}
	}
}

func TestMap002(t *testing.T) {
	m := New(1000)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100000; i++ {

		n := r.Int()
		m.Put(i, n)
		v, ok := m.Get(i)
		if !ok {
			t.Error()
		}
		if v != n {
			t.Error()
		}
	}
}

func TestMap003(t *testing.T) {
	m := New(10)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100; i++ {

		n := r.Int()
		m.Put(i, n)
		//m2[i] = n
		v, ok := m.Get(i)
		if !ok {
			t.Error()
		}
		if v != n {
			t.Error()
		}

		if m.Len() != 1 {
			t.Error()
		}

		success := m.Delete(i)
		if !success {
			t.Error(i)
		}
		_, ok = m.Get(i)
		if ok {
			t.Error(i)
		}

		if m.Len() != 0 {
			t.Error()
		}

	}

}

//
func TestMap004(t *testing.T) {
	m := New(1000)

	m2 := make(map[int]int)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 1000; i++ {

		n1 := r.Int()
		n2 := r.Int()
		m.Put(n1, n2)
		m2[n1] = n2
		v, ok := m.Get(n1)
		if !ok {
			t.Error()
		}
		if v != n2 {
			t.Error()
		}

		if m.Len() != len(m2) {
			t.Error()
		}

	}

	for k1, v1 := range m2 {

		v2, ok := m.Get(k1)
		if !ok {
			t.Error()
		}

		if v2 != v1 {
			t.Error()
		}

		success := m.Delete(k1)
		if !success {
			t.Error()
		}

		delete(m2, k1)

		if m.Len() != len(m2) {
			t.Error()
		}

	}

}

//
func TestMap005(t *testing.T) {
	m := New(1000)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100000; i++ {

		n1 := r.Int()
		//n2 := r.Int()
		success := m.Delete(n1)

		if success {
			t.Error()
		}

	}

}

//
func TestMap006(t *testing.T) {
	m := New(1)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100000; i++ {

		n1 := r.Int()
		//n2 := r.Int()
		_, ok := m.Get(n1)

		if ok {
			t.Error()
		}

	}

}

func TestMap007(t *testing.T) {
	m := New(1)

	n := m.hash(100)
	if n != 0 {
		t.Error()
	}

	m = New(100)

	n = m.hash(101)
	if n != 1 {
		t.Error()
	}

}

func TestMap008(t *testing.T) {
	m := New(100)

	m.Put(1, 1)
	m.Put(1, 2)

	n, ok := m.Get(1)
	if !ok {
		t.Error()
	}

	if n != 2 {
		t.Error()
	}

}

func TestMap009(t *testing.T) {
	m := New(100)

	m.Put(1, 1)
	m.Put(101, 101)

	success := m.Delete(201)
	if success {
		t.Error()
	}

}

func TestMap_Iterator001(t *testing.T) {
	m := New(100)

	m.Put(1, 1)
	m.Put(101, 101)

	iter := m.Iterator()
	for {
		kv := iter.Next()
		if kv == nil {
			break
		}
		//log.Println(kv)
	}
}

func TestMap_Iterator002(t *testing.T) {
	m := New(10000)

	m2 := map[int]int{}
	m3 := map[int]int{}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 10000; i++ {

		n1 := r.Int()
		n2 := r.Int()

		m.Put(n1, n2)

		m2[n1] = n2
	}

	iter := m.Iterator()
	for {
		kv := iter.Next()
		if kv == nil {
			break
		}
		//log.Println(kv)

		m3[kv.Key] = kv.Value

	}

	if !isSameKeyValueMap(m2, m3) {
		t.Error()
	}
}

func isSameKeyValueMap(m1 map[int]int, m2 map[int]int) bool {

	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {

		value2, ok := m2[k]
		if !ok {
			return false
		}

		if value2 != v {
			return false
		}
	}

	return true
}
