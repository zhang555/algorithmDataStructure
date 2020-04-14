package sort

func RadixSort(arr []int64) {

	var (
		bucket       = make([][]int64, 10)
		a      int64 = 1
		//b      float64 = 0.1
	)

	for i := 1; i < 19; i++ {
		a *= 10
		//b *= 10

		//log.Println(`arr1: `, arr)

		for index := range bucket {
			bucket[index] = nil
		}

		for _, v := range arr {
			bucketIndex := v % a
			//if int64(b) != 0 {
			bucketIndex /= a / 10

			bucket[bucketIndex] = append(bucket[bucketIndex], v)
		}

		//log.Println(`bucket: `, bucket)

		var (
			count int
		)
		for _, v := range bucket {

			//log.Println(`v: `, v)
			//if len(v) == 0 {
			//	continue
			//}
			//for vIndex := 0; vIndex <= len(v)/2; vIndex++ {
			//	v[vIndex], v[len(v)-1-vIndex] = v[len(v)-1-vIndex], v[vIndex]
			//}

			copy(arr[count:], v[:])

			//log.Println(`arr: `, arr)

			count += len(v)

			if count >= len(arr) {
				break
			}
		}

		//log.Println(`arr2: `, arr)
	}

}
