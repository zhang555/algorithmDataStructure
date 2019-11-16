package bloomfilter

import (
	"math/rand"
	"testing"
	"time"
)

func TestBloomFilter001(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	bf := New(1000, 3)

	for i := 0; i < 10000; i++ {
		n := r.Intn(10000000)

		bf.Set(n)

		if bf.MustNotExist(n) {
			t.Error()
		}

	}

}

func TestBloomFilter_Set(t *testing.T) {

	//r := rand.New(rand.NewSource(time.Now().Unix()))

	bf := New(1000, 3)

	bf.Set(1)
	bf.Set(2)
	bf.Set(3)

}
