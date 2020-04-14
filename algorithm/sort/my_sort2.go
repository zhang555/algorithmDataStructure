package sort

import (
	"runtime"
	"sort"
	"sync"
)

func MyMergeSort2(src []int64) {
	if len(src) < 1<<10 {
		ShellSort(src)
		return
	}

	var (
		wg     sync.WaitGroup
		cpuNum = runtime.NumCPU()
		size   = len(src) / cpuNum
	)

	wg.Add(cpuNum)

	for i := 0; i < cpuNum; i++ {
		go func(idx int) {
			defer wg.Done()
			//offset := idx * size

			if idx == cpuNum-1 {
				sort.Slice(src[idx*size:], func(x, y int) bool { return src[x] < src[y] })

			} else {
				sort.Slice(src[idx*size:idx*size+size], func(x, y int) bool { return src[x] < src[y] })

			}

		}(i)
	}

	wg.Wait()

	sort.Slice(src, func(x, y int) bool { return src[x] < src[y] })

}
