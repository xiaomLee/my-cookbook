/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 *
 * https://leetcode.cn/problems/number-of-enclaves/description/
 *
 * algorithms
 * Medium (61.93%)
 * Likes:    254
 * Dislikes: 0
 * Total Accepted:    70.2K
 * Total Submissions: 113.4K
 * Testcase Example:  '[[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]'
 *
 * 给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
 *
 * 一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
 *
 * 返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
 * 输出：3
 * 解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
 * 输出：0
 * 解释：所有 1 都在边界上或可以到达边界。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 500
 * grid[i][j] 的值为 0 或 1
 *
 *
 */

// @lc code=start
func numEnclaves(grid [][]int) int {
	// 从四周边界对岛屿进行淹没
	m := len(grid)
	n := len(grid[0])
	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}

	// 上下边界
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}

	// 左右边界
	for i := 0; i < n; i++ {
		dfs(grid, 0, i)
		dfs(grid, m-1, i)
	}

	// 统计
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

// @lc code=end

