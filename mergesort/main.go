package main

func main() {
	sortArray([]int{5, 2, 3, 1})
}

func sortArray(nums []int) []int {
	mergesort(nums, 0, len(nums)-1)
	return nums
}

// 将数组[l...r]一分为二，分别对左右数组进行排序，然后对排序好的数组进行归并
func mergesort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergesort(arr, l, mid)
	mergesort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

// [l...mid] [mid+1...r]
func merge(arr []int, l, mid, r int) {
	arr1 := arr[l : mid+1]
	arr2 := arr[mid+1 : r+1]
	newArr := make([]int, r-l+1)
	i := 0 // 当前遍历元素
	j := 0
	k := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			newArr[k] = arr2[j]
			j++
			k++
			continue
		}
		newArr[k] = arr1[i]
		k++
		i++
	}
	if i == len(arr1) {
		copy(newArr[k:], arr2[j:])
	}
	if j == len(arr2) {
		copy(newArr[k:], arr1[i:])
	}
	copy(arr[l:], newArr)
}
