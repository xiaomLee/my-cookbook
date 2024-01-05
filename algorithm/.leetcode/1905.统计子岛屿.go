/*
 * @lc app=leetcode.cn id=1905 lang=golang
 *
 * [1905] 统计子岛屿
 *
 * https://leetcode.cn/problems/count-sub-islands/description/
 *
 * algorithms
 * Medium (67.24%)
 * Likes:    112
 * Dislikes: 0
 * Total Accepted:    28.8K
 * Total Submissions: 42.8K
 * Testcase Example:  '[[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]]\n' +
  '[[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]'
 *
 * 给你两个 m x n 的二进制矩阵 grid1 和 grid2 ，它们只包含 0 （表示水域）和 1 （表示陆地）。一个 岛屿 是由 四个方向
 * （水平或者竖直）上相邻的 1 组成的区域。任何矩阵以外的区域都视为水域。
 *
 * 如果 grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，也就是说 grid2 中该岛屿的每一个格子都被 grid1
 * 中同一个岛屿完全包含，那么我们称 grid2 中的这个岛屿为 子岛屿 。
 *
 * 请你返回 grid2 中 子岛屿 的 数目 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]],
 * grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
 * 输出：3
 * 解释：如上图所示，左边为 grid1 ，右边为 grid2 。
 * grid2 中标红的 1 区域是子岛屿，总共有 3 个子岛屿。
 *
 *
 * 示例 2：
 *
 * 输入：grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]],
 * grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
 * 输出：2
 * 解释：如上图所示，左边为 grid1 ，右边为 grid2 。
 * grid2 中标红的 1 区域是子岛屿，总共有 2 个子岛屿。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid1.length == grid2.length
 * n == grid1[i].length == grid2[i].length
 * 1 <= m, n <= 500
 * grid1[i][j] 和 grid2[i][j] 都要么是 0 要么是 1 。
 *
 *
*/

// @lc code=start
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	// 思路 将grid2中的岛屿不存在grid1中的淹掉， 然后统计grid2岛屿的数量

	count := 0
	m := len(grid1)
	n := len(grid1[0])

	var dfs func(nums [][]int, i, j int)
	dfs = func(nums [][]int, i, j int) {
		if i < 0 || j < 0 || i >= len(nums) || j >= len(nums[0]) {
			return
		}
		if nums[i][j] == 0 {
			return
		}
		// 淹没当前岛屿：将所有相邻的 nums[i][j] == 1 => nums[i][j] = 0
		// 四个方向淹没
		nums[i][j] = 0
		dfs(nums, i, j+1)
		dfs(nums, i, j-1)
		dfs(nums, i+1, j)
		dfs(nums, i-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				// grid2 存在的岛屿， grid1不存在
				dfs(grid2, i, j)
			}
		}
	}

	// 统计剩余岛屿数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				// grid2 存在的岛屿， grid1不存在
				dfs(grid2, i, j)
				count++
			}
		}
	}
	return count
}

// @lc code=end

