package sort

/*
ShellSort 希尔排序

将key 设置成 len的一半， 对

*/
func ShellSort(src []int64) {
	key := len(src)
	for key > 0 {
		for i := key; i < len(src); i++ {
			j := i
			for j >= key && src[j] < src[j-key] {
				src[j], src[j-key] = src[j-key], src[j]
				j = j - key
			}
		}
		key = key / 2
	}
}
