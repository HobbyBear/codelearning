package quicksort

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 2))
}
