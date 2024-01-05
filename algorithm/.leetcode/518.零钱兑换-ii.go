/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	// dp[i] 表示组成金额 i 的组合数
	// 则 dp[i] += dp[i-coins[j]]
	// base case dp[0] = 1
	dp := make([]int, amount+1)
	for j := 0; j < len(coins); j++ {
		for i := coins; i <= amount; i++ {
			dp[i] += dp[i-coins[j]]
		}
	}
	return dp[amount]
}

// @lc code=end

