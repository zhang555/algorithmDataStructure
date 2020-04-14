package sort

import (
	"log"
	"math"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/pingcap/check"
)

var _ = check.Suite(&sortTestSuite{})

func TestT(t *testing.T) {
	//check.TestingT(t)

	conf := &check.RunConf{
		//Filter: `^TestSort`,
		//Filter: `^TestSort2$`,
		Filter: `TestSort$`,
		//Filter: `TestRadixSort$`,
		//Filter: `CountSort$`,
		//Filter: `TestHeapSort$`,
		//Filter: `sortTestSuite.TestSort2`,
	}

	check.RunAll(conf)
}

func TestName1(t *testing.T) {

	var (
		re  *regexp.Regexp
		err error
		b   bool
	)

	//re, err = regexp.Compile(`TestSort$`)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//b = re.MatchString(`TestSort`)
	//log.Println(b)
	//
	//b = re.MatchString(`TestSort2`)
	//log.Println(b)

	re, err = regexp.Compile(`(BenchmarkBubbleSort|BenchmarkGoSort2)$`)
	if err != nil {
		log.Fatal(err)
	}

	b = re.MatchString(`BenchmarkBubbleSort`)
	log.Println(b)

	b = re.MatchString(`BenchmarkGoSort2`)
	log.Println(b)

}

func prepare(src []int64) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		//src[i] = rand.Int63()
		src[i] = int64(rand.Uint32())
	}
}

type sortTestSuite struct{}

func (s *sortTestSuite) TestSort(c *check.C) {

	//lens := []int{1, 3, 5, 7, 11, 13, 17, 1
	//23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}
	//lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024}
	lens := []int{
		1 << 10,
		1 << 11,
	}
	//lens := []int{3}
	//lens := []int{30}
	//lens := []int{10}

	var testFuncs []func([]int64) = []func([]int64){
		//BubbleSort,
		//SelectionSort,
		//InsertSort,
		//ShellSort,
		//QuickSort,
		//HeapSort,
		//BucketSort,
		//RadixSort,
		//GoSort,
		//MyMergeSort,
		MyMergeSort2,
	}

	for _, testFunc := range testFuncs {
		for i := range lens {
			src := make([]int64, lens[i])
			expect := make([]int64, lens[i])
			prepare(src)
			copy(expect, src)
			testFunc(src)
			sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })

			//log.Println(src)
			//log.Println(expect)

			for i := 0; i < len(src); i++ {
				c.Assert(src[i], check.Equals, expect[i])
			}
		}
	}

}

func (s *sortTestSuite) TestSort2(c *check.C) {
	//lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024}
	//lens := []int{1 << 10}
	//lens := []int{3}
	//lens := []int{1024}
	//lens := []int{10}

	var testFuncs []func([]int64) []int64 = []func([]int64) []int64{
		//BubbleSort,
		//SelectionSort,
		//InsertSort,
		//ShellSort,
		MergeSort,
	}

	for _, testFunc := range testFuncs {
		for i := range lens {
			src := make([]int64, lens[i])
			expect := make([]int64, lens[i])
			prepare(src)
			copy(expect, src)
			result := testFunc(src)
			sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })

			//log.Println(src)
			//log.Println(expect)
			//log.Println(result)

			c.Log(expect)
			c.Log(result)

			for i := 0; i < len(src); i++ {
				c.Assert(result[i], check.Equals, expect[i])
			}
		}
	}

}

func (s *sortTestSuite) TestHeapSort(c *check.C) {

	//var (
	//	src []int64 = []int64{
	//		1, 2, 3,
	//	}
	//)

	src := make([]int64, 10)
	expect := make([]int64, 10)
	prepare2(src)
	copy(expect, src)
	HeapSort(src)
	sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })

	log.Println(src)
	log.Println(expect)

	for i := 0; i < len(src); i++ {
		c.Assert(src[i], check.Equals, expect[i])
	}

}

func (s *sortTestSuite) TestCountSort(c *check.C) {

	//var (
	//	src []int64 = []int64{
	//		1, 2, 3,
	//	}
	//)

	src := make([]int64, 10)
	expect := make([]int64, 10)
	prepare2(src)
	copy(expect, src)
	CountSort(src)
	sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })

	log.Println(src)
	log.Println(expect)

	for i := 0; i < len(src); i++ {
		c.Assert(src[i], check.Equals, expect[i])
	}

}

func (s *sortTestSuite) TestRadixSort(c *check.C) {

	//var (
	//	src []int64 = []int64{
	//		1, 2, 3,
	//	}
	//)
	l := 100000
	src := make([]int64, l)
	expect := make([]int64, l)
	prepare2(src)
	//log.Println(`src: `, src)

	copy(expect, src)

	RadixSort(src)
	sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })

	//log.Println(`expect: `, expect)
	//log.Println(`result: `, src)

	for i := 0; i < len(src); i++ {
		c.Assert(src[i], check.Equals, expect[i])
	}

}

func prepare2(src []int64) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		//src[i] = int64(rand.Intn(100000))
		//src[i] = int64(rand.Int63())
		src[i] = int64(rand.Uint32())
	}
}

func TestName3(t *testing.T) {

	log.Println(len(strconv.Itoa(math.MaxInt64)))

}
