package bitmap

import (
	"math/rand"
	"testing"
	"time"
)

func TestBitMap011(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	bitMap := BitMap{}

	for i := 0; i < 10000; i++ {
		n := r.Intn(10000000)

		bitMap.Set(uint(n))

		if bitMap.Get(uint(n)) != true {
			t.Error()
		}

		bitMap.Remove(uint(n))

		if bitMap.Get(uint(n)) != false {
			t.Error()
		}

	}

}
