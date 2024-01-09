/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N 皇后 II
 *
 * https://leetcode.cn/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (82.26%)
 * Likes:    496
 * Dislikes: 0
 * Total Accepted:    136.7K
 * Total Submissions: 166.2K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：2
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 9
 *
 *
 *
 *
 */

// @lc code=start
func totalNQueens(n int) int {
	var backtrack func(n int, row int, track [][]byte, res *int)
	backtrack = func(n, row int, track [][]byte, res *int) {
		if row == n {
			// fmt.Println(track)
			*res = *res + 1
			return
		}

		if track[row] == nil {
			track[row] = make([]byte, n)
			for col := 0; col < n; col++ {
				track[row][col] = '.'
			}
		}

		// 对于当前行的没一列进行选择
		for col := 0; col < n; col++ {
			// 对于当前的 track[row][col] = 'Q' 判断是否能攻击
			if canAttack(track, row, col) {
				continue
			}
			track[row][col] = 'Q'
			backtrack(n, row+1, track, res)
			track[row][col] = '.'
		}
	}

	track := make([][]byte, n)
	res := 0
	backtrack(n, 0, track, &res)
	return res
}

func canAttack(track [][]byte, row, col int) bool {
	// 同一列
	for i := 0; i < row; i++ {
		if track[i][col] == 'Q' {
			return true
		}
	}
	// 同一斜线 V 型 左斜上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if track[i][j] == 'Q' {
			return true
		}
	}
	// 同一斜线 V 型 右斜上方
	for i, j := row-1, col+1; i >= 0 && j < len(track); i, j = i-1, j+1 {
		if track[i][j] == 'Q' {
			return true
		}
	}
	return false
}

// @lc code=end

