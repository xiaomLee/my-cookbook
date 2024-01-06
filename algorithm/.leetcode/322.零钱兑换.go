/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// dp[i] 表示 能否从 coins 数组中凑成 i 金额
	// 则 range coins { dp[i] = min(dp[i], dp[i-coins[j]+1]) }
	// base case dp[0] = 0 dp[i] = amount+1

	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		if i == 0 {
			dp[i] = 0
			continue
		}
		dp[i] = amount + 1
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] < dp[i] {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// @lc code=end

