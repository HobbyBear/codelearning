package quicksort

import (
	"math/rand"
)

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

/**
	选定基点v， 让基点左边部分 < v 右边部分 > v ， 中间部分等于v
[l+1...lt] < v
[gt...r] > v
i 代表当前需要遍历的元素

快排每次分区排序后，能够让基点在数组中正确的位置上
*/

// 对[l..r]快排
func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	lt := l     // 小于v的右边界
	gt := r + 1 // 大于v的左边界
	i := l + 1  // 当前遍历的元素
	swap(arr, l, rand.Int()%(r-l+1)+l)
	v := arr[l]
	for i < gt {
		if v > arr[i] {
			tmp := arr[lt+1]
			arr[lt+1] = arr[i]
			arr[i] = tmp
			lt++
			i++
			continue
		}
		if v < arr[i] {
			gt--
			tmp := arr[gt]
			arr[gt] = arr[i]
			arr[i] = tmp
			continue
		}
		i++
	}
	tmp := arr[lt]
	arr[lt] = v
	arr[l] = tmp
	lt--
	quickSort(arr, l, lt)
	quickSort(arr, gt, r)
}
