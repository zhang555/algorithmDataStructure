package lru

import (
	"algorithmDataStructure/lru/lru1"
	"algorithmDataStructure/lru/lru2"
	"algorithmDataStructure/lru/lru3"
	"math/rand"
	"testing"
	"time"

	. "github.com/pingcap/check"
)

var _ = Suite(&testLRUSuite{})

type testLRUSuite struct {
	random *rand.Rand
}

func TestName(t *testing.T) {
	TestingT(t)
}

func (s *testLRUSuite) SetUpSuite(c *C) {
	s.random = rand.New(rand.NewSource(time.Now().Unix()))
}

func (s *testLRUSuite) TearDownSuite(c *C) {

}

func (s *testLRUSuite) SetUpTest(c *C) {

}

func (s *testLRUSuite) TearDownTest(c *C) {

}

func (s *testLRUSuite) Test(c *C) {
	for i := 0; i < 100; i++ {
		cap1 := s.random.Intn(100)
		lru11 := lru1.Constructor(cap1)
		lru22 := lru2.Constructor(cap1)
		lru33 := lru3.Constructor(cap1)

		for i := 0; i < 10000; i++ {
			switch s.random.Intn(2) {
			case 0:
				key := s.random.Intn(100)
				value := s.random.Intn(100)

				lru11.Put(key, value)
				lru22.Put(key, value)
				lru33.Put(key, value)
			case 1:
				key := s.random.Intn(100)

				num1 := lru11.Get(key)
				num2 := lru22.Get(key)
				num3 := lru33.Get(key)

				c.Assert(num1, Equals, num2)
				c.Assert(num1, Equals, num3)
			}
		}
	}

}
