package bitmap

type BitMap struct {
	slice []uint64
}

func (b *BitMap) Get(i uint) bool {
	var slicePosition = i >> 6
	var bitPosition uint64 = 1 << (i & (64 - 1))
	return b.slice[slicePosition]&bitPosition != 0
}

func (b *BitMap) Set(i uint) {
	b.extendIfNeed(i)
	b.slice[i>>6] |= 1 << (i & (64 - 1))
	return
}

func (b *BitMap) Remove(i uint) {
	b.slice[i>>6] &= ^(1 << (i & (64 - 1)))
}

func (b *BitMap) extendIfNeed(i uint) {
	var slicePosition = i >> 6
	if int(slicePosition) > len(b.slice)-1 {
		b.slice = append(b.slice, make([]uint64, int(slicePosition)-len(b.slice)+1)...)
	}
	return
}
