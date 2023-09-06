package bloomfilter

import (
	"codelearning/bitmap"
	"github.com/spaolacci/murmur3"
	"math"
)

type BloomFilter struct {
	bitset *bitmap.BitMap
	k      int
	m      int64
}

// m 代表布隆过滤器大小，也就是其中bitset的大小，k代表hash函数的个数
func New(m int64, k int) *BloomFilter {
	return &BloomFilter{
		bitset: bitmap.New(m),
		k:      k,
		m:      m,
	}
}

func (b *BloomFilter) Add(key string) {
	bitSetSize := b.m
	hashFunCount := b.k
	hash := murmur3.New128()
	hash.Write([]byte(key))
	hash1, hash2 := hash.Sum128()
	combinedHash := hash1
	for i := 0; i < hashFunCount; i++ {
		b.bitset.Set((uint64ToInt64(combinedHash) & math.MaxInt64) % bitSetSize)
		combinedHash += hash2
	}
}

func (b *BloomFilter) MightContain(key string) bool {
	bitSetSize := b.m
	hashFunCount := b.k
	hash := murmur3.New128()
	hash.Write([]byte(key))
	hash1, hash2 := hash.Sum128()
	combinedHash := hash1
	for i := 0; i < hashFunCount; i++ {
		if !b.bitset.Exits((uint64ToInt64(combinedHash) & math.MaxInt64) % bitSetSize) {
			return false
		}
		combinedHash += hash2
	}
	return true
}

func uint64ToInt64(num uint64) int64 {
	if num <= uint64(math.MaxInt64) {
		return int64(num)
	}
	return int64(num - uint64(math.MaxInt64) + 1)
}
