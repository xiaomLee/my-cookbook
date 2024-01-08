/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode.cn/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (69.90%)
 * Likes:    1621
 * Dislikes: 0
 * Total Accepted:    547.2K
 * Total Submissions: 782.8K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 200
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	// dp[i][j] 表示到达 grid[i][j] 的最小路径和
	// 状态转移: min(上、左) + grid[i][j]
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	// base case: dp[0][0] = grid[0][0] dp[0][j] = dp[0][j-1] + grid[0][j] ...
	// return dp[len(grid)-1][len(grid[0])-1]

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		if dp[i] == nil {
			dp[i] = make([]int, len(grid[0]))
		}
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
				continue
			}
			if i == 0 {
				dp[0][j] = dp[0][j-1] + grid[0][j]
			} else if j == 0 {
				dp[i][0] = dp[i-1][0] + grid[i][0]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

// @lc code=end

