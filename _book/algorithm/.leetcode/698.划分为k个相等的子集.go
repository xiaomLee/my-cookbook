/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 *
 * https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/description/
 *
 * algorithms
 * Medium (41.89%)
 * Likes:    997
 * Dislikes: 0
 * Total Accepted:    110.1K
 * Total Submissions: 262.8K
 * Testcase Example:  '[4,3,2,3,5,2,1]\n4'
 *
 * 给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
 * 输出： True
 * 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,2,3,4], k = 3
 * 输出: false
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= len(nums) <= 16
 * 0 < nums[i] < 10000
 * 每个元素的频率在 [1,4] 范围内
 *
 *
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	visited := 0
	memo := make(map[int]bool)
	sort.Ints(nums)
	return backtrack(nums, k, target, 0, 0, visited, memo)
}

func backtrack(nums []int, k int, target int, pos, sum int, visited int, memo map[int]bool) bool {
	if k == 0 {
		return true
	}
	if sum == target {
		res := backtrack(nums, k-1, target, 0, 0, visited, memo)
		memo[visited] = res
		return res
	}

	if res, ok := memo[visited]; ok {
		return res
	}

	for i := pos; i < len(nums); i++ {
		if ((visited >> i) & 1) == 1 {
			continue
		}
		if sum+nums[i] > target {
			continue
		}
		sum += nums[i]
		visited |= 1 << i
		if backtrack(nums, k, target, pos+1, sum, visited, memo) {
			return true
		}
		sum -= nums[i]
		visited ^= 1 << i
	}
	return false
}

// @lc code=end

