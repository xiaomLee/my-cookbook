/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (59.99%)
 * Likes:    1643
 * Dislikes: 0
 * Total Accepted:    304.9K
 * Total Submissions: 508.3K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：prices = [3,3,5,0,0,3,1,4]
 * 输出：6
 * 解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
 * 随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
 *
 * 示例 2：
 *
 *
 * 输入：prices = [1,2,3,4,5]
 * 输出：4
 * 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4
 * 。
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 *
 *
 * 示例 3：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
 *
 * 示例 4：
 *
 *
 * 输入：prices = [1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp[i][j][k] 表示第 i 天在最大进行 j 次交易的情况下分别持有、不持有股票时的最大收益
	// 完成交易以买入时间计算，买入时 j-1
	// 则 有如下状态转移
	// 1. 第 i 天最多可进行 j 笔交易 且不持有股票: 前一日也不持仓 dp[i-1][k][0]， 前一日持仓今日卖出 dp[i-1][j][1] + prices[i]
	// 2. 第 i 天最多可进行 j 笔交易 且持仓：前一日持仓 dp[i-1][j][1] , 前一日不持仓(前一日最大可用-1) 今日买入 dp[i-1][j-1][0] - prices[i]
	// base case dp[0][j][0] = 0, dp[0][j][1] = -price[i]
	// return dp[len(prices)-1][2][0]

	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		if dp[i] == nil {
			dp[i] = make([][]int, 3)
		}
		for j := 0; j < 3; j++ {
			if dp[i][j] == nil {
				dp[i][j] = make([]int, 2)
			}
			if i == 0 {
				dp[0][j][0] = 0
				dp[0][j][1] = -prices[0] // 第一天持仓
				continue
			}
			if j == 0 { // 无法交易
				dp[i][0][0] = 0
				dp[i][0][1] = 0
				continue
			}
			// 第 i 天最多可进行 j 笔交易 且不持有股票
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			// 第 i 天最多可进行 j 笔交易 且持仓
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][2][0]
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}
	sort.Ints(nums)
	return nums[len(nums)-1]
}

// @lc code=end

