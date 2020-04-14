package hashmap

import (
	"math/rand"
	"testing"
	"time"
)

//func BenchmarkMap001(b *testing.B) {
//
//	m := New(1)
//
//	r := rand.New(rand.NewSource(time.Now().Unix()))
//
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < 100; j++ {
//			n1 := r.Int()
//			n2 := r.Int()
//			m.Put(n1, n2)
//		}
//	}
//}

func BenchmarkMap003(b *testing.B) {

	m := New(100000)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < b.N; i++ {
		//for j := 0; j < 100; j++ {
		n1 := r.Int()
		n2 := r.Int()
		m.Put(n1, n2)
		//}
	}
}

func BenchmarkMap002(b *testing.B) {

	m2 := make(map[int]int)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	//n1 := r.Int()
	//n2 := r.Int()

	for i := 0; i < b.N; i++ {

		//m2[n1] = n2

		//for j := 0; j < 100; j++ {
		n1 := r.Int()
		n2 := r.Int()
		m2[n1] = n2

		//}

	}
}
