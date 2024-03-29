/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 *
 * https://leetcode.cn/problems/coin-change-ii/description/
 *
 * algorithms
 * Medium (70.88%)
 * Likes:    1216
 * Dislikes: 0
 * Total Accepted:    282.3K
 * Total Submissions: 398K
 * Testcase Example:  '5\n[1,2,5]'
 *
 * 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
 * 
 * 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
 * 
 * 假设每一种面额的硬币有无限个。 
 * 
 * 题目数据保证结果符合 32 位带符号整数。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：amount = 5, coins = [1, 2, 5]
 * 输出：4
 * 解释：有四种方式可以凑成总金额：
 * 5=5
 * 5=2+2+1
 * 5=2+1+1+1
 * 5=1+1+1+1+1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：amount = 3, coins = [2]
 * 输出：0
 * 解释：只用面额 2 的硬币不能凑成总金额 3 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：amount = 10, coins = [10] 
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * coins 中的所有值 互不相同
 * 0 
 * 
 * 
 */

// @lc code=start
func change(amount int, coins []int) int {
	// dp[i] 表示 凑出 金额 i 的 硬币组合数
	// i 金额的组合数可由所有的 i-coins[j] 累加而来
	// 则 dp[i] += dp[i-coins[:]]
	// base case dp[0] = 1 
	dp:=make([]int, amount+1)
	dp[0] = 1
	for i:=0; i<len(coins); i++ {
		for j:=1; j<=amount+1; j++ {
			if coins[i] < j {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
// @lc code=end

