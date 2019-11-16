package bloomfilter

import (
	"algorithmDataStructure/bitmap"
	"crypto/md5"
	"strconv"
)

type BloomFilter struct {
	m      uint
	k      uint
	bitmap *bitmap.BitMap
}

//
func New(m uint, k uint) *BloomFilter {
	if m < 1 {
		panic(`m can not less than 1`)
	}
	if k < 1 {
		panic(`k can not less than 1`)
	}

	return &BloomFilter{
		m:      m,
		k:      k,
		bitmap: bitmap.New(),
	}
}

func (bf *BloomFilter) Set(v int) {
	results := bf.hash(v)
	for _, v := range results {
		bf.bitmap.Set(v)
	}
	return
}

func (bf *BloomFilter) Get(v int) bool {
	results := bf.hash(v)
	for _, v := range results {
		if !bf.bitmap.Get(v) {
			return false
		}
	}
	return true
}

func (bf BloomFilter) MayExist(v int) bool {
	return bf.Get(v)
}

func (bf BloomFilter) MustNotExist(v int) bool {
	return !bf.Get(v)
}

func (bf BloomFilter) hash(v int) []uint {
	var (
		results []uint
		result  uint
	)

	var i uint = 0
	for ; i < bf.k; i++ {
		string1 := strconv.Itoa(v) + strconv.Itoa(int(i))
		byteArray := md5.Sum([]byte(string1))
		for _, v := range byteArray {
			result += uint(v)
		}
		results = append(results, result%bf.m)
	}
	return results
}
