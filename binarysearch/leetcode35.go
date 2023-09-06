package binarysearch

func searchInsert(nums []int, target int) int {
	return searchInsertHelper(nums, 0, len(nums)-1, target)
}

// 在[l...r] 中寻找元素插入的索引值，即找到数组中第一个大于等于该值的索引位置
func searchInsertHelper(arr []int, l, r int, num int) int {
	// 定义终止条件
	if l > r {
		return l
	}
	mid := (r - l) / 2
	if arr[mid] == num {
		return mid
	}
	if arr[mid] >= num {
		// 左半部分寻找
		return searchHelper(arr, l, mid-1, num)
	}
	// arr[mid] < num 右半部分寻找
	return searchHelper(arr, mid+1, r, num)
}
