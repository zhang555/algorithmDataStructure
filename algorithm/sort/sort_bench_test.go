package sort

import "testing"

// go test -bench=.
//go test -run=XXX -bench=.

/*

go test   -test.v -test.bench . -test.run ^$

go test    -test.bench  ^BenchmarkBubbleSort$

go test -test.bench     ^(BenchmarkBubbleSort|BenchmarkGoSort2)$
go test -test.bench    (BenchmarkBubbleSort|BenchmarkGoSort2)$

(hello|world)
*/

func benchmark(b *testing.B, l int, sort func(src []int64)) {
	src := make([]int64, l)
	prepare(src)
	for i := 0; i < b.N; i++ {
		sort(src)
	}
}

func benchmark11(b *testing.B, sort func(src []int64), l ...int) {
	for _, v := range l {
		src := make([]int64, v)
		prepare(src)
		for i := 0; i < b.N; i++ {
			sort(src)
		}
	}

}

func benchmark2(b *testing.B, l int, sort func(src []int64) []int64) {
	src := make([]int64, l)
	prepare(src)
	for i := 0; i < b.N; i++ {
		sort(src)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	benchmark(b, 100, BubbleSort)
}

func BenchmarkGoSort(b *testing.B) {
	benchmark(b, 100, GoSort)
}

func BenchmarkHeapSort(b *testing.B) {
	benchmark(b, 100, HeapSort)
}

func BenchmarkInsertSort(b *testing.B) {
	benchmark(b, 100, InsertSort)
}

func BenchmarkMergeSort(b *testing.B) {
	benchmark2(b, 100, MergeSort)
}

func BenchmarkQuickSort(b *testing.B) {
	benchmark(b, 100, QuickSort)
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmark(b, 100, SelectionSort)
}

func BenchmarkShellSort(b *testing.B) {
	benchmark(b, 100, ShellSort)
}

func BenchmarkBucketSort(b *testing.B) {
	benchmark(b, 100, BucketSort)
}

func BenchmarkRadixSort(b *testing.B) {
	benchmark(b, 100, RadixSort)
}

//func BenchmarkCountSort(b *testing.B) {
//	benchmark2(b, 100, CountSort)
//}
