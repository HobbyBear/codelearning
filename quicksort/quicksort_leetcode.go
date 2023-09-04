package quicksort

import (
	"math/rand"
	"time"
)

/*
*

215. 数组中的第K个最大元素
中等
2.3K
相关企业
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

提示：

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/
func findKthLargest(nums []int, k int) int {
	return partition(nums, k, 0, len(nums)-1)
}

// 大于 v     小于v
// v [l+1..lt]     [gt...r] 进行分区，判断 分区后的lt+1 和k的大小
func partition(nums []int, k int, l, r int) int {
	lt := l
	rand.Seed(time.Now().Unix())
	swap(nums, l, rand.Intn(r-l+1)+l)
	v := nums[l]
	gt := r + 1
	i := l + 1 // 当前遍历元素
	for i < gt {
		if nums[i] < v {
			gt--
			swap(nums, gt, i)
			continue
		}
		if nums[i] > v {
			lt++
			swap(nums, i, lt)
			i++
			continue
		}
		i++
	}
	swap(nums, l, lt)
	lt--
	// lt+ 1的元素处于正确位置lt+1
	if lt+1 == k-1 {
		return nums[lt+1]
	}
	if lt+1 < k-1 {
		offset := lt + 2
		if gt >= k-1 {
			offset = k - 1
		}
		return partition(nums, k, offset, r)
	}
	return partition(nums, k, l, lt)
}
