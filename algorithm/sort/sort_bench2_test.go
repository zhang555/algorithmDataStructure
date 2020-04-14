package sort

import "testing"

var (
	len1 = 16000000
)

/*

go test -test.bench     ^(BenchmarkBubbleSort|BenchmarkGoSort2)$
go test -test.bench    (BenchmarkMyMergeSort|BenchmarkMyMergeSort2|BenchmarkMyMergeSort3|BenchmarkGoSort2)$

*/

func BenchmarkMyMergeSort(b *testing.B) {
	benchmark(b, len1, MyMergeSort)
}

func BenchmarkMyMergeSort2(b *testing.B) {
	benchmark(b, len1, MyMergeSort2)
}

func BenchmarkMyMergeSort3(b *testing.B) {
	benchmark(b, len1, MyMergeSort3)
}

func BenchmarkGoSort2(b *testing.B) {
	benchmark(b, len1, GoSort)
}

//func BenchmarkMyMergeSort2(b *testing.B) {
//	benchmark11(b, MyMergeSort2,
//		1<<10,
//		1<<11,
//		1<<13,
//	)
//}

//func BenchmarkGoSort22(b *testing.B) {
//	benchmark11(b, GoSort,
//		1<<10,
//		1<<11,
//		1<<13,
//	)
//}
