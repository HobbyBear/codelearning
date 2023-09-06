package heap

import (
	"fmt"
	"testing"
)

func TestHeapify(t *testing.T) {
	arr := []int{10, 100, 90, 1, 4, 20, 8}
	h := Heapify(arr)
	fmt.Println(h.arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(h.Pop())
	}
}
