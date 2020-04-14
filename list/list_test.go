package lru

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/pingcap/check"
)

var _ = Suite(&testListSuite{})

type testListSuite struct {
	random *rand.Rand
}

func TestName(t *testing.T) {
	TestingT(t)
}

func (s *testListSuite) SetUpSuite(c *C) {
	s.random = rand.New(rand.NewSource(time.Now().Unix()))
}

func (s *testListSuite) TearDownSuite(c *C) {

}

func (s *testListSuite) SetUpTest(c *C) {

}

func (s *testListSuite) TearDownTest(c *C) {

}

func (s *testListSuite) Test(c *C) {
	//for i := 0; i < 100; i++ {
	//	//cap1 := s.random.Intn(100)
	//	list := list.New()
	//	list1 := doublylinkedlist.New()
	//	//lru22 := lru2.Constructor(cap1)
	//	//lru33 := lru3.Constructor(cap1)
	//
	//	for i := 0; i < 10000; i++ {
	//		switch s.random.Intn(2) {
	//		case 0:
	//			key := s.random.Intn(100)
	//			value := s.random.Intn(100)
	//
	//			lru11.Put(key, value)
	//			lru22.Put(key, value)
	//			lru33.Put(key, value)
	//		case 1:
	//			key := s.random.Intn(100)
	//
	//			num1 := lru11.Get(key)
	//			num2 := lru22.Get(key)
	//			num3 := lru33.Get(key)
	//
	//			c.Assert(num1, Equals, num2)
	//			c.Assert(num1, Equals, num3)
	//		}
	//	}
	//}

}
