package binarysearch

func Search(arr []int, num int) int {
	return searchHelper(arr, 0, len(arr)-1, num)
}

// 在[l..r]中寻找num
func searchHelper(arr []int, l, r int, num int) int {
	// 定义终止条件
	if l > r {
		return -1
	}
	mid := (r - l) / 2
	if arr[mid] == num {
		return mid
	}
	if arr[mid] > num {
		// 左半部分寻找
		return searchHelper(arr, l, mid-1, num)
	}
	// arr[mid] < num 右半部分寻找
	return searchHelper(arr, mid+1, r, num)
}
