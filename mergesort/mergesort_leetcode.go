package main

/*
剑指 Offer 51. 数组中的逆序对
困难
1.1K
相关企业
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

示例 1:

输入: [7,5,6,4]
输出: 5

限制：

0 <= 数组长度 <= 50000
*/

func reversePairs(nums []int) int {
	var cnt = 0
	mergesort_copy(nums, 0, len(nums)-1, &cnt)
	return cnt
}

// 将数组[l...r]一分为二，分别对左右数组进行排序，然后对排序好的数组进行归并
func mergesort_copy(arr []int, l, r int, cnt *int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergesort_copy(arr, l, mid, cnt)
	mergesort_copy(arr, mid+1, r, cnt)
	mergeCopy(arr, l, mid, r, cnt)
}

// [l...mid] [mid+1...r]

func mergeCopy(arr []int, l, mid, r int, cnt *int) {
	arr1 := arr[l : mid+1]
	arr2 := arr[mid+1 : r+1]
	newArr := make([]int, r-l+1)
	i := 0 // 当前遍历元素
	j := 0
	k := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			newArr[k] = arr2[j]
			*cnt += len(arr1) - i
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
