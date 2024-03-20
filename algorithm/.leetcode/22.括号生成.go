/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.52%)
 * Likes:    3481
 * Dislikes: 0
 * Total Accepted:    777.9K
 * Total Submissions: 1M
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	// 共有2*n个字符， 左括号 右括号分别为 n 个
	// 递归， 每次递归分别减少一个左括号 或 右括号
	// 前序位置判断当前组合是否合法
	// left == 0 && right == 0 找到合法组合
	// left < 0 || right < 0 || left > right 非法
	var backtrack func(left, right int, track []byte, res *[]string)
	backtrack = func(left, right int, track []byte, res *[]string) {
		if right == 0 && left == 0 {
			*res = append(*res, string(track))
			return
		}
		if left < 0 || right < 0 || left > right {
			return
		}

		// 尝试添加左括号
		track = append(track, '(')
		backtrack(left-1, right, track, res)
		track = track[:len(track)-1]

		// 尝试添加右括号
		track = append(track, ')')
		backtrack(left, right-1, track, res)
		track = track[:len(track)-1]
	}
	track := make([]byte, 0)
	res := make([]string, 0)
	backtrack(n, n, track, &res)
	return res
}

// @lc code=end

