/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode.cn/problems/coin-change/description/
 *
 * algorithms
 * Medium (47.18%)
 * Likes:    2685
 * Dislikes: 0
 * Total Accepted:    730.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
 * 
 * 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 * 
 * 你可以认为每种硬币的数量是无限的。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3 
 * 解释：11 = 5 + 5 + 1
 * 
 * 示例 2：
 * 
 * 
 * 输入：coins = [2], amount = 3
 * 输出：-1
 * 
 * 示例 3：
 * 
 * 
 * 输入：coins = [1], amount = 0
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 * 
 * 
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// dp[i] 表示从 coins 中凑出 i 金额的 最小硬币数
	// 则 对于每一个硬币都可以从中选择，dp[i] = 最小的那个选择 + 1
	// dp[i] = min(dp[i], dp[i-coins[:]] + 1)
	// base case dp[i] = amount+1 dp[0] = 0
	
	dp := make([]int, amount+1)
	dp[0] = 0
	for i:=1; i<=amount; i++ {
		dp[i] = amount + 1
		for j:=0; j<len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] + 1 < dp[i] {
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

