package sort

//BubbleSort 冒泡排序， 相邻的，如果顺序不对，就交换，
func BubbleSort(src []int64) {
	for i := 1; i < len(src); i++ {
		for j := 0; j < len(src)-i; j++ {
			if src[j] > src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
}
