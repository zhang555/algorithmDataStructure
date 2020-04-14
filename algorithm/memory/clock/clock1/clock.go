package clock1

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

		if len(this.m) == this.cap {
			for i := 0; i < this.cap; i++ {
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

func (this *Clock) PositionIncrease() {
	this.position++
	if this.position >= this.cap {
		this.position = 0
	}
}

func (this *Clock) Debug() []KV {
	return this.list
}
