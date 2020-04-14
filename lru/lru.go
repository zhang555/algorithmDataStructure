package lru

type LRUer interface {
	Get(int) int
	Put(int, int)
}
