/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (37.90%)
 * Likes:    7008
 * Dislikes: 0
 * Total Accepted:    1.6M
 * Total Submissions: 4.2M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 * 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	// dp[i][j] 表示 s[i:j] 是否是回文串
	// 则 状态转移
	// 1. s[i] != s[j] dp[i][j] = false
	// 2. s[i] == s[j] dp[i][j] = dp[i+1][j-1]
	// base case: dp[0][0] = true dp[i][j] = true && i==j
	dp := make([][]bool, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		if dp[i] == nil {
			dp[i] = make([]bool, len(s))
		}
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = true
				continue
			}
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			}
		}
	}

	res := ""

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}

	return res
}

func longestPalindrome(s string) string {
	res := ""
	huiwen := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		if huiwen[i] == nil {
			huiwen[i] = make([]bool, len(s)+1)
		}
		for j := 0; j < i; j++ {
			huiwen[j][i] = isHuiWen(s[j:i])
			if huiwen[j][i] && i-j > len(res) {
				res = s[j:i]
			}
		}
	}
	return res
}

func isHuiWen(s string) (res bool) {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end

