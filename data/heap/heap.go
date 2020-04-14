package heap

type Heap struct {
	array []int
}

func New(array []int) *Heap {
	HeapAll(array)
	h := Heap{array: array}
	return &h
}

func (h *Heap) Insert(v int) {
	h.array = append(h.array, v)
	i := len(h.array) - 1
	for {
		if i-1 < 0 {
			return
		}
		p := (i - 1) / 2
		if h.array[i] < h.array[p] {
			h.array[i], h.array[p] = h.array[p], h.array[i]
		}
		i = p
	}
}

func (h *Heap) DeleteMin() {
	if len(h.array) == 0 {
		return
	}
	if len(h.array) == 1 {
		h.array = h.array[:len(h.array)-1]
		return
	}
	i := 0
	last := h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	for {
		l := 2*i + 1
		r := 2*i + 2
		if l >= len(h.array) {
			h.array[i] = last
			return
		}
		if r >= len(h.array) {
			if last <= h.array[l] {
				h.array[i] = last
				return
			} else {
				h.array[i] = h.array[l]
				i = l
				continue
			}
		}

		if last <= h.array[l] && last <= h.array[r] {
			h.array[i] = last
			return
		}

		n := r
		if h.array[l] < h.array[r] {
			n = l
		}
		h.array[i] = h.array[n]
		i = n
	}
}

func (h *Heap) FindMin() int {
	if len(h.array) != 0 {
		return h.array[0]
	}
	return 0
}
