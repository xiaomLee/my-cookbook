/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode.cn/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (59.94%)
 * Likes:    2391
 * Dislikes: 0
 * Total Accepted:    748.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1
 * grid[i][j] 的值为 '0' 或 '1'
 *
 *
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	count := 0
	m := len(grid)
	n := len(grid[0])

	var dfs func(nums [][]byte, i, j int)
	dfs = func(nums [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= len(nums) || j >= len(nums[0]) {
			return
		}
		if nums[i][j] == '0' {
			return
		}
		// 淹没当前岛屿：将所有相邻的 nums[i][j] == '1' => nums[i][j] = '0'
		// 四个方向淹没
		nums[i][j] = '0'
		dfs(nums, i, j+1)
		dfs(nums, i, j-1)
		dfs(nums, i+1, j)
		dfs(nums, i-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

// @lc code=end

