package sort

import "math"

//SelectionSort 选择排序  遍历后面n个 ，找到最小的，插入到n-1
func SelectionSort(src []int64) {
	for j := 0; j < len(src)-1; j++ {
		var (
			min      int64 = math.MaxInt64
			minIndex int
		)
		for i := j; i < len(src); i++ {
			if src[i] < min {
				min = src[i]
				minIndex = i
			}
		}

		if minIndex != 0 {
			src[j], src[minIndex] = src[minIndex], src[j]
		}
	}
}
