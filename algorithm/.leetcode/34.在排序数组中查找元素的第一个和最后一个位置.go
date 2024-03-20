/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (42.88%)
 * Likes:    2607
 * Dislikes: 0
 * Total Accepted:    908.5K
 * Total Submissions: 2.1M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
 * 
 * 如果数组中不存在目标值 target，返回 [-1, -1]。
 * 
 * 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [5,7,7,8,8,10], target = 8
 * 输出：[3,4]
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,7,7,8,8,10], target = 6
 * 输出：[-1,-1]
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [], target = 0
 * 输出：[-1,-1]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * nums 是一个非递减数组
 * -10^9 <= target <= 10^9
 * 
 * 
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
        return []int{-1, -1}
    }

	res := make([]int, 2)

	// 分别二分查找左边界、右边界

	// 左边界 right 左移
	left, right := 0, len(nums)-1
	for left + 1 < right {
		mid := left  + (right - left)/2
		if nums[mid] < target {
			left = mid
		}else if nums[mid] > target {
			right = mid
		}else {
			// nums[mid] == target
			right = mid
		}
	}
	if nums[left] == target {
		res[0] = left
	}else if nums[right] == target {
		res[0] = right
	}else {
		res[0] = -1
		res[1] = -1
	}
	if res[0] == -1 {
		return res
	}

	// 右边界
	left, right = 0, len(nums)-1
	for left + 1 < right {
		mid := left  + (right - left)/2
		if nums[mid] < target {
			left = mid
		}else if nums[mid] > target {
			right = mid
		}else {
			// nums[mid] == target
			left = mid
		}
	}
	if nums[right] == target {
		res[1] = right
	}else if nums[left] == target {
		res[1] = left
	}else {
		res[1] = -1
	}

	return res
}
// @lc code=end

