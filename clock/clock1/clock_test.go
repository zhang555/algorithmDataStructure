package clock1_test

import (
	"algorithmDataStructure/clock/clock1"
	"math/rand"
	"testing"
	"time"

	. "github.com/pingcap/check"
)

var _ = Suite(&testSuite{})

type testSuite struct {
	random *rand.Rand
}

func TestName(t *testing.T) {
	TestingT(t)
}

func (s *testSuite) SetUpSuite(c *C) {
	s.random = rand.New(rand.NewSource(time.Now().Unix()))
}

func (s *testSuite) TearDownSuite(c *C) {

}

func (s *testSuite) SetUpTest(c *C) {

}

func (s *testSuite) TearDownTest(c *C) {

}

func (s *testSuite) Test(c *C) {
	for i := 0; i < 100; i++ {
		//cap1 := s.random.Intn(100)
		//lru11 := clock1.Constructor(cap1)
		////lru22 := lru2.Constructor(cap1)
		////lru33 := lru3.Constructor(cap1)
		//
		//for i := 0; i < 10000; i++ {
		//	switch s.random.Intn(2) {
		//	case 0:
		//		key := s.random.Intn(100)
		//		value := s.random.Intn(100)
		//
		//		lru11.Put(key, value)
		//		lru22.Put(key, value)
		//		lru33.Put(key, value)
		//	case 1:
		//		key := s.random.Intn(100)
		//
		//		num1 := lru11.Get(key)
		//		num2 := lru22.Get(key)
		//		num3 := lru33.Get(key)
		//
		//		c.Assert(num1, Equals, num2)
		//		c.Assert(num1, Equals, num3)
		//	}
		//}
	}

}

func (s *testSuite) Test2(c *C) {

	lru11 := clock1.Constructor(3)

	lru11.Put(2, 2)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 2, Value: 2, IsUse: true},
		{Key: 0, Value: 0, IsUse: false},
		{Key: 0, Value: 0, IsUse: false},
	})

	lru11.Put(3, 3)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 2, Value: 2, IsUse: true},
		{Key: 3, Value: 3, IsUse: true},
		{Key: 0, Value: 0, IsUse: false},
	})

	lru11.Put(2, 2)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 2, Value: 2, IsUse: true},
		{Key: 3, Value: 3, IsUse: true},
		{Key: 0, Value: 0, IsUse: false},
	})

	lru11.Put(1, 1)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 2, Value: 2, IsUse: true},
		{Key: 3, Value: 3, IsUse: true},
		{Key: 1, Value: 1, IsUse: true},
	})

	lru11.Put(5, 5)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 5, Value: 5, IsUse: true},
		{Key: 3, Value: 3, IsUse: false},
		{Key: 1, Value: 1, IsUse: false},
	})

	lru11.Put(2, 2)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 5, Value: 5, IsUse: true},
		{Key: 2, Value: 2, IsUse: true},
		{Key: 1, Value: 1, IsUse: false},
	})

	lru11.Put(4, 4)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 5, Value: 5, IsUse: true},
		{Key: 2, Value: 2, IsUse: true},
		{Key: 4, Value: 4, IsUse: true},
	})

	lru11.Put(5, 5)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 5, Value: 5, IsUse: true},
		{Key: 2, Value: 2, IsUse: true},
		{Key: 4, Value: 4, IsUse: true},
	})

	lru11.Put(3, 3)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 3, Value: 3, IsUse: true},
		{Key: 2, Value: 2, IsUse: false},
		{Key: 4, Value: 4, IsUse: false},
	})

	lru11.Put(2, 2)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 3, Value: 3, IsUse: true},
		{Key: 2, Value: 2, IsUse: true},
		{Key: 4, Value: 4, IsUse: false},
	})

	lru11.Put(5, 5)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 3, Value: 3, IsUse: true},
		{Key: 2, Value: 2, IsUse: false},
		{Key: 5, Value: 5, IsUse: true},
	})

	lru11.Put(2, 2)
	Judge(c, lru11.Debug(), []clock1.KV{
		{Key: 3, Value: 3, IsUse: true},
		{Key: 2, Value: 2, IsUse: true},
		{Key: 5, Value: 5, IsUse: true},
	})

}

func Judge(c *C, kv1 []clock1.KV, kv2 []clock1.KV) {

	c.Assert(len(kv1), Equals, len(kv2))

	for i := range kv1 {

		c.Assert(kv1[i].Key, Equals, kv2[i].Key)
		c.Assert(kv1[i].Value, Equals, kv2[i].Value)
		c.Assert(kv1[i].IsUse, Equals, kv2[i].IsUse)
	}

}
