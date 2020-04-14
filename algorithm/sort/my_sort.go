package sort

import (
	"runtime"
	"sort"
	"sync"
)

type sliceHeader struct {
	data []int64
	idx  int
}

func MyMergeSort(src []int64) {

	var (
		wg     sync.WaitGroup
		cpuNum = runtime.NumCPU()
		size   = len(src) / cpuNum
	)

	if len(src) < cpuNum {
		cpuNum = len(src)
	}

	// 收集各个 goroutine 排好序的数组
	resultChan := make(chan sliceHeader, cpuNum)

	wg.Add(cpuNum)

	for i := 0; i < cpuNum; i++ {
		go func(idx int) {
			defer wg.Done()
			offset := idx * size

			// copy to avoid race
			dst := make([]int64, size)
			copy(dst, src[offset:offset+size])

			// 最后一个数组需要把那些剩余的元素一起管理起来
			if idx == cpuNum-1 {
				for j := offset + size; j < len(src); j++ {
					dst = append(dst, src[j])
				}
			}

			sort.Slice(dst, func(x, y int) bool { return dst[x] < dst[y] })
			resultChan <- sliceHeader{data: dst, idx: 0}
		}(i)
	}

	wg.Wait()
	close(resultChan)

	var mulArr []sliceHeader

	// 基本的并发结果收集套路
	for arr := range resultChan {
		mulArr = append(mulArr, arr)
	}

	buildHeap(mulArr)

	// 结果需要拷贝回原来的数组
	resultArr := src[:0]
	// 从堆顶的数组取元素
	// 取完之后需要判断堆顶数组是否已经消耗完毕
	// 消耗完毕则 popHeap，未消耗完毕，则修改 idx，重平衡 heap
	for len(mulArr) > 0 && mulArr[0].idx < len(mulArr[0].data) {
		resultArr = append(resultArr, mulArr[0].data[mulArr[0].idx])
		mulArr[0].idx++
		if mulArr[0].idx > len(mulArr[0].data)-1 {
			mulArr = popHeap(mulArr)
		} else {
			adjustHeap(mulArr, 0, len(mulArr)-1)
		}
	}
}

func buildHeap(arr []sliceHeader) {
	for i := len(arr) / 2; i >= 0; i-- {
		adjustHeap(arr, i, len(arr)-1)
	}
}

// popHeap 必须 return arr
// 否则局部对 slice len 的修改无法反应在原来的变量上
func popHeap(arr []sliceHeader) []sliceHeader {
	lastIdx := len(arr) - 1
	arr[0], arr[lastIdx] = arr[lastIdx], arr[0]
	arr = arr[:lastIdx]
	adjustHeap(arr, 0, len(arr)-1)
	return arr
}

// 小顶堆 adjust
func adjustHeap(arr []sliceHeader, headIdx, tailIdx int) {
	// 长度 <=1 ，不需要 adjust
	if len(arr) <= 1 {
		return
	}

	cur := headIdx
	// todo optimize
	// too ugly
	for {
		// 在当前元素，左右子节点元素中，找到最小的元素
		// 如果该元素和当前元素索引不同，则交换之，继续向下走
		// 否则说明堆已经调整完毕
		lIdx, rIdx, maxIdx, min := 2*cur+1, 2*cur+2, cur, arr[cur].data[arr[cur].idx]
		if lIdx <= tailIdx && arr[lIdx].data[arr[lIdx].idx] < min {
			maxIdx, min = lIdx, arr[lIdx].data[arr[lIdx].idx]
		}
		if rIdx <= tailIdx && arr[rIdx].data[arr[rIdx].idx] < min {
			maxIdx, min = rIdx, arr[rIdx].data[arr[rIdx].idx]
		}
		if maxIdx == cur {
			break
		}

		arr[cur], arr[maxIdx] = arr[maxIdx], arr[cur]
		cur = maxIdx
	}
}
