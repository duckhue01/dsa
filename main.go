package main

import (
	"hash"
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	bits     []bool
	k        uint64
	hashFunc hash.Hash64
}

func NewBloomFilter(m uint64, k uint64) *BloomFilter {
	return &BloomFilter{
		bits:     make([]bool, m),
		k:        k,
		hashFunc: fnv.New64(),
	}
}

func (b *BloomFilter) add(item []byte) {
	for i := uint64(0); i < b.k; i++ {
		b.hashFunc.Reset()
		b.hashFunc.Write(item)
		hashVal := b.hashFunc.Sum64()
		index := hashVal % uint64(len(b.bits))
		b.bits[index] = true
	}
}

func (b *BloomFilter) contains(item []byte) bool {
	for i := uint64(0); i < b.k; i++ {
		b.hashFunc.Reset()
		b.hashFunc.Write(item)
		hashVal := b.hashFunc.Sum64()
		index := hashVal % uint64(len(b.bits))
		if !b.bits[index] {
			return false
		}
	}
	return true
}

func optimalParameters(n uint64, p float64) (uint64, uint64) {
	m := math.Ceil((-1 * float64(n) * math.Log(p)) / math.Pow(math.Log(2), 2))
	k := math.Ceil((math.Log(2) * m) / float64(n))
	return uint64(m), uint64(k)
}
func main() {
	b := NewBloomFilter(3, 3)

}
