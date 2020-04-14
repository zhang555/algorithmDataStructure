package sort

import "math"

func BucketSort(src []int64) {

	var (
		bucketNum int       = 100
		bucket    [][]int64 = make([][]int64, bucketNum)
		ret       []int64
	)

	for _, v := range src {
		bucket[int(v)/(math.MaxInt64/bucketNum)] = append(bucket[int(v)/(math.MaxInt64/bucketNum)], v)
	}

	for _, v := range bucket {
		InsertSort(v)
		ret = append(ret, v...)
	}
	copy(src, ret)
}
