/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// dp[i] 表示以 nums[i] 为结尾的最长递增子序列的长度, 注意是子序列(可不连续)不是子串
	// 则 dp[i] = for j<i { dp[i] = max (nums[j]<nums[i]? dp[j]+1 : dp[i]  })
	// base case dp[0] = 1
	// return max(dp[:])
	dp := make([]int, len(nums))
	dp[0] = 1
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// @lc code=end

