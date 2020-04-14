package map1

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/pingcap/check"
)

var _ = Suite(&testSuite{})

type testSuite struct {
	random *rand.Rand
}

/*

go test -check.f "Test2"
go test -check.f "Test33"
go test -check.f "TestRehash"


*/
func TestName(t *testing.T) {

	TestingT(t)
}

func TestName111111(t *testing.T) {
	conf := &RunConf{
		//Filter: `Test1$`,
		//Filter: `Test3$`,
		//Filter: `Test33$`,
		Filter: `TestRehash$`,
	}
	RunAll(conf)

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

	for i := 0; i < 1; i++ {
		//cap1 := s.random.Intn(1000)
		m := New(10000)
		m1 := make(map[int]int)

		for i := 0; i < 100; i++ {
			switch s.random.Intn(2) {
			case 0:
				key := s.random.Intn(100)
				value := s.random.Intn(100)

				m.Put(key, value)
				c.Log(key, value)
				m1[key] = value

			case 1:
				key := s.random.Intn(100)

				num1, ok1 := m.Get(key)
				c.Log(`get `, key, num1, ok1)
				num2, ok2 := m1[key]

				c.Assert(num1, Equals, num2)
				c.Assert(ok1, Equals, ok2)
			}
		}
	}

}

func (s *testSuite) Test2(c *C) {
	m := New(10)
	for i := 0; i < 10; i++ {
		key := 1 + i*10
		value := key

		m.Put(key, value)

		num1, ok1 := m.Get(key)
		c.Assert(num1, Equals, value)
		c.Assert(ok1, IsTrue)

	}
}

func (s *testSuite) Test3(c *C) {

	length := 10
	m := New(length)
	for i := 0; i < length; i++ {

		key := 10
		value := s.random.Intn(10)

		m.Put(key, value)

		num1, ok1 := m.Get(key)
		c.Assert(num1, Equals, value)
		c.Assert(ok1, IsTrue)
	}
}

func (s *testSuite) Test4(c *C) {

	length := 10
	m := New(length)
	for i := 0; i < length; i++ {

		key := i
		value := s.random.Intn(100)

		m.Put(key, value)

		num1, ok1 := m.Get(key)
		c.Assert(num1, Equals, value)
		c.Assert(ok1, IsTrue)

	}

}

/*
go test -check.f "Test33"

*/
func (s *testSuite) Test33(c *C) {
	for i := 0; i < 1000; i++ {
		m := New(1)
		m1 := make(map[int]int)

		for i := 0; i < 10000; i++ {
			switch s.random.Intn(3) {
			case 0:
				key := s.random.Intn(10000)
				value := s.random.Intn(10000)

				m.Put(key, value)
				m1[key] = value

			case 1:
				key := s.random.Intn(10000)

				num1, ok1 := m.Get(key)
				num2, ok2 := m1[key]

				//c.Assert(num1, Equals, num2)
				c.Check(num1, Equals, num2)
				//c.Assert(ok1, Equals, ok2)
				c.Check(ok1, Equals, ok2)
			case 2:
				//key := s.random.Intn(10000)
				//m.Delete(key)
				//delete(m1, key)
			}
		}

	}

	//c.Log(111)
	//log.Println(len(c.GetTestLog()))

}

/*
go test -check.f "TestRehash"

*/

func (s *testSuite) TestRehash(c *C) {
	for i := 0; i < 1000; i++ {
		m := New(1)
		//m1 := make(map[int]int)

		for i := 0; i < 10; i++ {
			key := s.random.Intn(10000)
			value := s.random.Intn(10000)

			m.Put(key, value)
		}
	}
}
