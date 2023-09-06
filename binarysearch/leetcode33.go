package binarysearch

func search(nums []int, target int) int {
	return searchhelper(nums, 0, len(nums)-1, target)
}

// 查找[start,end] 范围内有没有target的值
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2]
func searchhelper(nums []int, start int, end int, target int) int {
	if start > end {
		return -1
	}
	mid := (end-start)/2 + start
	if nums[mid] == target {
		return mid
	}
	// 左边数组,  target 如果在start end范围内，正常二分
	if nums[mid] >= nums[0] {
		if nums[mid] > target && target >= nums[start] {
			return searchhelper(nums, start, mid-1, target)
		}
		return searchhelper(nums, mid+1, end, target)
	}
	// 右边数组
	if nums[mid] < target && target <= nums[end] {
		return searchhelper(nums, mid+1, end, target)
	}
	return searchhelper(nums, start, mid-1, target)
}
