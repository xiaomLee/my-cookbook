/*
 * @lc app=leetcode.cn id=1254 lang=golang
 *
 * [1254] 统计封闭岛屿的数目
 *
 * https://leetcode.cn/problems/number-of-closed-islands/description/
 *
 * algorithms
 * Medium (63.95%)
 * Likes:    278
 * Dislikes: 0
 * Total Accepted:    61.5K
 * Total Submissions: 96.1K
 * Testcase Example:  '[[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]'
 *
 * 二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
 *
 * 请返回 封闭岛屿 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：grid =
 * [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
 * 输出：2
 * 解释：
 * 灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[1,1,1,1,1,1,1],
 * [1,0,0,0,0,0,1],
 * [1,0,1,1,1,0,1],
 * [1,0,1,0,1,0,1],
 * [1,0,1,1,1,0,1],
 * [1,0,0,0,0,0,1],
 * ⁠            [1,1,1,1,1,1,1]]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= grid.length, grid[0].length <= 100
 * 0 <= grid[i][j] <=1
 *
 *
 */

// @lc code=start
func closedIsland(grid [][]int) int {
	// 从四周边界对岛屿进行淹没
	m := len(grid)
	n := len(grid[0])
	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 1 {
			return
		}
		grid[i][j] = 1
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
			if grid[i][j] == 0 {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

// @lc code=end

