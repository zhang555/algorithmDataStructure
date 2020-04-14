package sort

import "sort"

/*
727109541
81834120
1487231636
*/
func GoSort(src []int64) {
	sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
}
