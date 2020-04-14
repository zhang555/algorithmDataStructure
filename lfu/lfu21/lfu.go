package lfu21

import (
	"log"
)

type LFUCache struct {
	keyMap  map[int]*Node
	freqMap map[int]*List2
	cap     int
	minFreq int
}

type KV struct {
	Key   int
	Value int
	Times int
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.LstdFlags)
}

func Constructor(capacity int) LFUCache {
	lfu := LFUCache{
		keyMap:  map[int]*Node{},
		freqMap: map[int]*List2{},
		cap:     capacity,
		minFreq: 1,
	}
	lfu.freqMap[1] = New()
	lfu.freqMap[2] = New()
	return lfu
}

func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}

	v, ok := this.keyMap[key]
	if !ok {
		return -1
	}

	l1 := this.freqMap[v.Value.(*KV).Times]
	l1.Remove(v)

	if l1.Len() == 0 && this.minFreq == v.Value.(*KV).Times {
		this.minFreq = v.Value.(*KV).Times + 1
	}

	v.Value.(*KV).Times++

	l1 = this.freqMap[v.Value.(*KV).Times]
	if l1 == nil {
		l1 = New()
		this.freqMap[v.Value.(*KV).Times] = l1
	}
	l1.PushFront(v)
	return v.Value.(*KV).Value
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	v, ok := this.keyMap[key]
	if !ok {
		if len(this.keyMap) == this.cap {
			//删除 最少使用 最不常使用的
			l1 := this.freqMap[this.minFreq]
			if l1 == nil {
				panic(` l1 == nil `)
			}
			element := l1.Back()
			//log.Println(l1.Len())
			l1.Remove(element)
			//log.Println(l1.Len())
			delete(this.keyMap, element.Value.(*KV).Key)
		}

		//添加新元素
		l1 := this.freqMap[1]
		if l1 == nil {
			l1 = New()
			this.freqMap[1] = l1
		}

		n := &Node{Value: &KV{Key: key, Value: value, Times: 1}}
		l1.PushFront(n)
		this.minFreq = 1
		return
	}

	//老元素 ，设置新的值
	l1 := this.freqMap[v.Value.(*KV).Times]
	l1.Remove(v)

	if l1.Len() == 0 && this.minFreq == v.Value.(*KV).Times {
		this.minFreq = v.Value.(*KV).Times + 1
	}

	v.Value.(*KV).Times++
	v.Value.(*KV).Value = value

	l1 = this.freqMap[v.Value.(*KV).Times]
	if l1 == nil {
		l1 = New()
		this.freqMap[v.Value.(*KV).Times] = l1
	}

	l1.PushFront(v)

	return
}
