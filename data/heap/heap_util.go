package heap

//判断是不是小顶堆
func IsSmallHeap(array []int) bool {
	for i := 0; i <= len(array)/2; i++ {
		l := 2*i + 1
		r := 2*i + 2
		if l >= len(array) {
			return true
		} else if array[i] > array[l] {
			return false
		}
		if r >= len(array) {
			return true
		} else if array[i] > array[r] {
			return false
		}
	}
	return true
}

//HeapOnce 将 i 这个位置上的元素 保证 比 子节点小， 如果替换了某个子节点，就递归
func HeapOnce(array []int, i int, end int) {

	l := 2*i + 1
	if l >= end {
		return
	}
	r := 2*i + 2
	t := l
	if r < end && array[l] >= array[r] {
		t = r
	}
	if array[i] <= array[t] {
		return
	}
	array[i], array[t] = array[t], array[i]
	HeapOnce(array, t, end)
}

func HeapAll(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		HeapOnce(array, i, len(array))
	}

}
