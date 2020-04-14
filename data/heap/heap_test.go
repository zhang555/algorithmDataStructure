package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/pingcap/check"
)

func TestName(t *testing.T) {
	check.TestingT(t)
}

func TestName2(t *testing.T) {
	conf := &check.RunConf{
		//Filter: `Test1$`,
		//Filter: `Test3$`,
		//Filter: `TestHeap$`,
		//Filter: `TestHeap_DeleteMin$`,
		//Filter: `TestHeapAll2$`,
	}
	check.RunAll(conf)
}

func prepare(src []int) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		//src[i] = rand.Int63()
		src[i] = int(rand.Uint32())
	}
}

var _ = check.Suite(&testSuite{})

type testSuite struct{}

func (s *testSuite) SetUpSuite(c *check.C) {
	rand.Seed(time.Now().Unix())
}

func (s *testSuite) TearDownSuite(c *check.C) {

}

func (s *testSuite) SetUpTest(c *check.C) {

}

func (s *testSuite) TearDownTest(c *check.C) {

}

func (s *testSuite) TestHeapAll(c *check.C) {

	src := make([]int, 100000)
	prepare(src)

	HeapAll(src)

	c.Assert(IsSmallHeap(src), check.Equals, true)

}

func (s *testSuite) TestHeapAll2(c *check.C) {

	//src := make([]int, 100000)
	//prepare(src)

	var src []int = []int{3, 2, 1}

	HeapAll(src)

	c.Assert(IsSmallHeap(src), check.Equals, true)

}

func (s *testSuite) TestIsHeap(c *check.C) {

	var src []int = []int{3, 1, 2}
	c.Assert(IsSmallHeap(src), check.Equals, false)

	src = []int{3, 1, 3}
	c.Assert(IsSmallHeap(src), check.Equals, false)

	src = []int{3, 1, 4}
	c.Assert(IsSmallHeap(src), check.Equals, false)

	src = []int{4, 1, 4, 1}
	c.Assert(IsSmallHeap(src), check.Equals, false)

	src = []int{4, 1, 4, 2}
	c.Assert(IsSmallHeap(src), check.Equals, false)

	src = []int{1, 2, 3, 4}
	c.Assert(IsSmallHeap(src), check.Equals, true)

}

func (s *testSuite) TestHeap(c *check.C) {

	h := New(nil)
	for i := 0; i < 10000; i++ {
		h.Insert(rand.Int())

		c.Assert(IsSmallHeap(h.array), check.Equals, true)
		c.Assert(len(h.array), check.Equals, i+1)

	}

}

func (s *testSuite) TestHeap_DeleteMin(c *check.C) {

	h := New(nil)

	times := 10000

	for i := 0; i < times; i++ {
		h.Insert(i)
		c.Assert(IsSmallHeap(h.array), check.Equals, true)
		c.Assert(len(h.array), check.Equals, i+1)
	}

	for i := 0; i < times; i++ {
		c.Assert(h.FindMin(), check.Equals, i)
		h.DeleteMin()
		c.Log(i)
		c.Assert(IsSmallHeap(h.array), check.Equals, true)
	}

	//h := New(nil)
	//h.Insert(1)
	//h.Insert(2)
	//h.Insert(3)
	//
	//c.Assert(IsSmallHeap(h.array), check.Equals, true)
	//
	//c.Assert(h.FindMin(), check.Equals, 1)
	//h.DeleteMin()
	//c.Assert(h.FindMin(), check.Equals, 2)
	//h.DeleteMin()
	//c.Assert(h.FindMin(), check.Equals, 3)
	//h.DeleteMin()

}
