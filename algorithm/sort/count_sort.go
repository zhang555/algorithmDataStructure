package sort

import "math"

func CountSort(arr []int64) []int64 {
	var (
		m int64 = math.MinInt64
	)

	for _, v := range arr {
		if v >= m {
			m = v
		}
	}

	var (
		bucket   = make([]int64, m+1)
		arrIndex int
	)

	for _, v := range arr {
		bucket[v]++
	}

	for i, v := range bucket {
		for v > 0 {
			arr[arrIndex] = int64(i)
			v--
			arrIndex++
		}
	}

	return arr

}
