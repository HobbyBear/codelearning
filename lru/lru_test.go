package lru

import (
	"fmt"
	"testing"
)

func TestLRU_Add(t *testing.T) {
	cache := New(4)
	cache.Add(3)
	fmt.Println("0", cache)
	cache.Add(4)
	fmt.Println("1", cache)
	cache.Add(4)
	fmt.Println("2", cache)
	cache.Add(6)
	cache.Add(7)
	fmt.Println("3", cache)
	cache.Add(8)
	fmt.Println("4", cache)

}
