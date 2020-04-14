package hashmap

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/pingcap/check"
)

func TestMap000(t *testing.T) {
	m := New(1000)

	m.Put(1, 1)

	v, ok := m.Get(1)
	if !ok {
		t.Error()
	}
	if v != 1 {
		t.Error()
	}

}

func TestMap001(t *testing.T) {
	m := New(1000)

	for i := 0; i < 100; i++ {
		m.Put(i, i)
		v, ok := m.Get(i)
		if !ok {
			t.Error()
		}
		if v != i {
			t.Error()
		}
	}
}

func TestMap002(t *testing.T) {
	m := New(1000)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100000; i++ {

		n := r.Int()
		m.Put(i, n)
		v, ok := m.Get(i)
		if !ok {
			t.Error()
		}
		if v != n {
			t.Error()
		}
	}
}

func TestMap003(t *testing.T) {
	m := New(1000)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 1000000; i++ {

		n := r.Int()
		m.Put(i, n)
		//m2[i] = n
		v, ok := m.Get(i)
		if !ok {
			t.Error()
		}
		if v != n {
			t.Error()
		}

		if m.Len() != 1 {
			t.Error()
		}

		success := m.Delete(i)
		if !success {
			t.Error()
		}
		_, ok = m.Get(i)
		if ok {
			t.Error()
		}

		if m.Len() != 0 {
			t.Error()
		}

	}

}

//
func TestMap004(t *testing.T) {
	m := New(1000)

	m2 := make(map[int]int)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100000; i++ {

		n1 := r.Int()
		n2 := r.Int()
		m.Put(n1, n2)
		m2[n1] = n2
		v, ok := m.Get(n1)
		if !ok {
			t.Error()
		}
		if v != n2 {
			t.Error()
		}

		if m.Len() != len(m2) {
			t.Error()
		}

	}

	for k1, v1 := range m2 {

		v2, ok := m.Get(k1)
		if !ok {
			t.Error()
		}

		if v2 != v1 {
			t.Error()
		}

		success := m.Delete(k1)
		if !success {
			t.Error()
		}

		delete(m2, k1)

		if m.Len() != len(m2) {
			t.Error()
		}

	}

}

//
func TestMap005(t *testing.T) {
	m := New(1000)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100000; i++ {

		n1 := r.Int()
		//n2 := r.Int()
		success := m.Delete(n1)

		if success {
			t.Error()
		}

	}

}

//
func TestMap006(t *testing.T) {
	m := New(1)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 100000; i++ {

		n1 := r.Int()
		//n2 := r.Int()
		_, ok := m.Get(n1)

		if ok {
			t.Error()
		}

	}

}

func TestMap007(t *testing.T) {
	m := New(1)

	n := m.hash(100)
	if n != 0 {
		t.Error()
	}

	m = New(100)

	n = m.hash(101)
	if n != 1 {
		t.Error()
	}

}

func TestMap008(t *testing.T) {
	m := New(100)

	m.Put(1, 1)
	m.Put(1, 2)

	n, ok := m.Get(1)
	if !ok {
		t.Error()
	}

	if n != 2 {
		t.Error()
	}

}

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

func TestMap009(t *testing.T) {
	m := New(100)

	m.Put(1, 1)
	m.Put(101, 101)

	success := m.Delete(201)
	if success {
		t.Error()
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

				c.Assert(num1, Equals, num2)
				//c.Check(num1, Equals, num2)
				c.Assert(ok1, Equals, ok2)
				//c.Check(ok1, Equals, ok2)
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
