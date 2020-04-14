package lru2

import (
	"algorithmDataStructure/lru"
	"testing"
)

func TestLRU1(t *testing.T) {
	testHelper(t, &lru.Test1)
}

func TestLRU2(t *testing.T) {
	testHelper(t, &lru.Test2)
}

func TestLRU3(t *testing.T) {
	testHelper(t, &lru.Test3)
}

func TestLRU4(t *testing.T) {
	testHelper(t, &lru.Test4)
}

func TestLRU5(t *testing.T) {
	testHelper(t, &lru.Test5)
}

func TestLRU6(t *testing.T) {
	testHelper(t, &lru.Test6)
}

func TestLRU7(t *testing.T) {
	testHelper(t, &lru.Test7)
}

func testHelper(t *testing.T, test1 *lru.LRUTest) {
	l1 := Constructor(test1.Cap)

	for _, v := range test1.Params {
		if v.IsGet {
			if l1.Get(v.P1) != v.P2 {
				t.Error()
			}
		} else {
			l1.Put(v.P1, v.P2)
		}
	}

}
