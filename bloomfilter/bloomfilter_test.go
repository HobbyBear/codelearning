package bloomfilter

import (
	"fmt"
	"math"
	"testing"
)

func TestBloomFilter_Add(t *testing.T) {
	bf := New(10, 3)
	bf.Add("test2")
	fmt.Println(bf.MightContain("test2"))
}

func TestNew(t *testing.T) {
	var maxuint uint64 = math.MaxUint64
	fmt.Println(maxuint)
	fmt.Println(int64(maxuint))
	fmt.Println(uint64ToInt64(maxuint) & math.MaxInt64)
}
