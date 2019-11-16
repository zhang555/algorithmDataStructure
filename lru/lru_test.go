package lru

import "testing"

func TestLRU(t *testing.T) {

	lru := Constructor(1)

	if lru.Get(1) != -1 {
		t.Error()
	}

	lru.Put(1, 1)

	if lru.Get(1) != 1 {
		t.Error()
	}

	lru.Put(2, 3)

	if lru.Get(1) != -1 {
		t.Error()
	}

	if lru.Get(2) != 3 {
		t.Error()
	}

}

func TestLRU01(t *testing.T) {

	lru := Constructor(1)

	lru.Put(1, 1)

	lru.Put(2, 3)

	if lru.Get(2) != 3 {
		t.Error()
	}

}

func TestLRU02(t *testing.T) {

	lru := Constructor(1)

	lru.Put(2, 1)
	if lru.Get(2) != 1 {
		t.Error()
	}

	lru.Put(3, 2)

	if lru.Get(2) != -1 {
		t.Error()
	}

	if lru.Get(3) != 2 {
		t.Error()
	}

}

func TestLRU2(t *testing.T) {

	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)

	if lru.Get(1) != 1 {
		t.Error()
	}

	lru.Put(3, 3)

	if lru.Get(1) != 1 {
		t.Error()
	}

	if lru.Get(2) != -1 {
		t.Error()
	}

	if lru.Get(3) != 3 {
		t.Error()
	}

}

func TestLRU3(t *testing.T) {

	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)

	if lru.Get(1) != 1 {
		t.Error()
	}

	lru.Put(3, 3)

	if lru.Get(2) != -1 {
		t.Error()
	}

	lru.Put(4, 4)

	if lru.Get(1) != -1 {
		t.Error()
	}

	if lru.Get(3) != 3 {
		t.Error()
	}

	if lru.Get(4) != 4 {
		t.Error()
	}

}

func TestLRU4(t *testing.T) {

	lru := Constructor(2)

	lru.Put(2, 1)
	lru.Put(2, 2)

	if lru.Get(2) != 2 {
		t.Error()
	}

	lru.Put(1, 1)
	lru.Put(4, 1)

	if lru.Get(2) != -1 {
		t.Error()
	}

}
