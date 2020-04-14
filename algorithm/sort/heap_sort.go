package sort

/*
堆排序
*/
func HeapSort(array []int64) {
	m := len(array)
	s := m / 2

	//从小到大排序， 先构造最大堆 , 从堆的角度来看， 从下往上处理， 将大值放上面 ，小值放下面
	for i := s; i > -1; i-- {
		heap(array, i, m-1)
	}

	//array[0]放的是最大的值， 将最大的值 放到list最后面
	for i := m - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]

		//交换之后，除去最后面的元素 对前面的元素维护最大堆
		heap(array, 0, i-1)
	}
}

//heap 将指定节点的值 交换为较大的值， 并对被换的值也递归处理
func heap(array []int64, i, end int) {
	l := 2*i + 1
	if l > end {
		return
	}
	n := l
	r := 2*i + 2
	if r <= end && array[r] > array[l] {
		n = r
	}
	if array[i] > array[n] {
		return
	}
	array[n], array[i] = array[i], array[n]
	heap(array, n, end)
}
