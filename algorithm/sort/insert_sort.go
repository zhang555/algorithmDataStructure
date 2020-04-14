package sort

//InsertSort 插入排序 ， 保证前n个是 排序的， 把 n+1 插入进去
//func InsertSort(src []int64) {
//	for i := 1; i < len(src); i++ {
//		var (
//			tmp    int64
//			change bool
//		)
//		tmp = src[i]
//		for j := 0; j <= i; j++ {
//			if change {
//				src[j], tmp = tmp, src[j]
//			} else {
//				if tmp < src[j] {
//					change = true
//					src[j], tmp = tmp, src[j]
//				}
//			}
//		}
//	}
//}

func InsertSort(src []int64) {
	for i := 1; i < len(src); i++ {
		for j := i; j >= 1 && src[j] < src[j-1]; j-- {
			src[j], src[j-1] = src[j-1], src[j]
		}
	}
}
