package sort

/*
快速排序
选取基准值， 比基准值大的，放到右边， 小的放到左边，
然后将前面和后面 递归
*/
func QuickSort(src []int64) {
	if len(src) <= 1 {
		return
	}
	var (
		head int
		tail int   = len(src) - 1
		mid  int64 = src[0]
	)
	for i := 1; i <= tail; {
		if src[i] > mid {
			src[i], src[tail] = src[tail], src[i]
			tail--
		} else {
			src[i], src[head] = src[head], src[i]
			head++
			i++
		}
	}
	QuickSort(src[:head])
	QuickSort(src[head+1:])
}

//func QuickSort(data []int64) {
//	if len(data) <= 1 {
//		return
//	}
//	mid := data[0]
//	head, tail := 0, len(data)-1
//	for i := 1; i <= tail; {
//		if data[i] > mid {
//			data[i], data[tail] = data[tail], data[i]
//			tail--
//		} else {
//			data[i], data[head] = data[head], data[i]
//			head++
//			i++
//		}
//	}
//	QuickSort(data[:head])
//	QuickSort(data[head+1:])
//}
