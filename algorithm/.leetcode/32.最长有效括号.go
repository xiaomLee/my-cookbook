/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode.cn/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (37.57%)
 * Likes:    2453
 * Dislikes: 0
 * Total Accepted:    424.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '"(()"'
 *
 * 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "(()"
 * 输出：2
 * 解释：最长有效括号子串是 "()"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = ")()())"
 * 输出：4
 * 解释：最长有效括号子串是 "()()"
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = ""
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * s[i] 为 '(' 或 ')'
 * 
 * 
 * 
 * 
 */

// @lc code=start
func longestValidParentheses(s string) int {
	// dp[i] 表示以 s[i] 为结尾的最长有效括号长度
	// 1 s[i] == ')' 
	// 1.1 s[i-1] == '(' 则 dp[i] = dp[i-2] + 2
	// 1.2. s[i] == ')' s[i-1] == ')' if s[i-dp[i-1]-1] == '(' dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2] else dp[i] = 0
	// 1.3  s[i] == '(' dp[i] = 0
	// base case: dp[0] = 0

	dp := make([]int, len(s))
	res := 0

	for i:=1; i<len(s); i++ {
		if s[i] == ')' && s[i-1] == '(' {
			if i > 2 {
				dp[i] = dp[i-2] + 2
			}else {
				dp[i] = 2
			}
		}else if s[i] == ')' && s[i-1] == ')' {
			if dp[i-1] > 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 > 0 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				}else {
					dp[i] = dp[i-1] + 2
				}
			}else {
				dp[i] = 0
			}
		}else {
			dp[i] = 0
		}

		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
// @lc code=end

