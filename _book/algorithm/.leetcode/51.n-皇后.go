/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 *
 * https://leetcode.cn/problems/n-queens/description/
 *
 * algorithms
 * Hard (73.95%)
 * Likes:    2002
 * Dislikes: 0
 * Total Accepted:    357.4K
 * Total Submissions: 483.3K
 * Testcase Example:  '4'
 *
 * 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
 *
 *
 *
 * 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4
 * 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[["Q"]]
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
func solveNQueens(n int) [][]string {
	var backtrack func(n int, row int, track [][]byte, res *[][]string)
	backtrack = func(n, row int, track [][]byte, res *[][]string) {
		if row == n {
			// fmt.Println(track)
			ans := make([]string, 0)
			for _, bytes := range track {
				// fmt.Println(string(bytes))
				ans = append(ans, string(bytes))
			}
			*res = append(*res, ans)
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
	res := make([][]string, 0)
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

