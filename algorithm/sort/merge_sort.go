package sort

/*
归并排序

将数据分成两段，从两端中 选取小的 按顺序 放到新数组里面，

*/
func MergeSort(src []int64) []int64 {
	if len(src) <= 1 {
		return src
	}
	a := MergeSort(src[:len(src)/2])
	b := MergeSort(src[len(src)/2:])
	return merge(a, b)
}

func merge(a []int64, b []int64) []int64 {
	ret := make([]int64, len(a)+len(b))
	var (
		i int
		j int
		k int
	)

	for {
		if a[i] < b[j] {
			ret[k] = a[i]
			i++
			k++
		} else {
			ret[k] = b[j]
			j++
			k++
		}

		if i >= len(a) {
			copy(ret[k:], b[j:])
			break
		}

		if j >= len(b) {
			copy(ret[k:], a[i:])
			break
		}

		if k == len(a)+len(b) {
			break
		}
	}
	return ret
}
