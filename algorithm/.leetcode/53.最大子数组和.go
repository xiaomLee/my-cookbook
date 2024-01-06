/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// dp[i] 表示以 nums[i] 为结尾的最大子数组和
	// 则 dp[i] = dp[i-1]> 0 ? dp[i-1] + num[i] : nums[i]
	// base case dp[0] = nums[0]
	// return max(dp[:])
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// @lc code=end

