package clock1

/*
时钟置换算法
每一个页框关联一个附加位， 首次装入内存时，附加位为1 ， 如果被访问到， 置1 。

需要替换时，查找isUse为0的， 如果遇到isUse为1的置0，然后继续查找下一个，
如果查找到了最后一个页框，从最后一个页框跳转到第一个页框 进行查找
如果全部查找一遍 没有isUse为0的页框，就替换当前页框。
*/
type Clock struct {
	m        map[int]int
	list     []KV
	cap      int
	position int
}

type KV struct {
	Key   int
	Value int
	IsUse bool
}

func Constructor(capacity int) Clock {
	return Clock{
		m:    make(map[int]int, capacity),
		list: make([]KV, capacity),
		cap:  capacity,
	}
}

func (this *Clock) Get(key int) int {
	index, ok := this.m[key]
	if !ok {
		return -1
	}
	this.list[index].IsUse = true
	return this.list[index].Value
}

func (this *Clock) Put(key int, value int) {
	index, ok := this.m[key]
	if ok {
		this.list[index].IsUse = true
		this.list[index].Key = key
		this.list[index].Value = value
		return
	} else {
		if len(this.m) == 0 {
			this.list[this.position].IsUse = true
			this.list[this.position].Key = key
			this.list[this.position].Value = value

			this.m[key] = this.position
			this.PositionIncrease()
			return
		}

		//找到一个需要丢弃的页框
		if len(this.m) == this.cap {
			for i := 0; i < this.cap; i++ {

				//IsUse == false 替换
				if !this.list[this.position].IsUse {
					delete(this.m, this.list[this.position].Key)
					this.m[key] = this.position
					this.list[this.position].Key = key
					this.list[this.position].Value = value
					this.list[this.position].IsUse = true
					this.PositionIncrease()
					return
				} else {
					this.list[this.position].IsUse = false
				}

				this.PositionIncrease()
			}

			delete(this.m, this.list[this.position].Key)
			this.m[key] = this.position
			this.list[this.position].Key = key
			this.list[this.position].Value = value
			this.list[this.position].IsUse = true
			this.PositionIncrease()
			return

		}

		//
		this.m[key] = this.position
		this.list[this.position].Key = key
		this.list[this.position].Value = value
		this.list[this.position].IsUse = true
		this.PositionIncrease()
		return

	}
}

//PositionIncrease 页框位置加一，如果超出，就置0
func (this *Clock) PositionIncrease() {
	this.position++
	if this.position >= this.cap {
		this.position = 0
	}
}

func (this *Clock) Debug() []KV {
	return this.list
}
