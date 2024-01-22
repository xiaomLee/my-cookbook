/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode.cn/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (37.56%)
 * Likes:    2433
 * Dislikes: 0
 * Total Accepted:    417.6K
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
	// 则 
	// 1. 若 s[i] == '(' 以 s[i] 为结尾的最长有效括号为必定为 0 ： dp[i] = 0
	// 2. 若 s[i] == ')' 需根据 s[i-1] 做状态转移
	// 2.1 若 s[i-1] == '(' 则 s[i-1] s[i] 已能组成有效括号 dp[i] = dp[i-2] + 2
	// 2.2 若 s[i-1] == ')' 则 需判断 dp[i-1] 
	// 2.2.1 若 dp[i-1] == 0 则 dp[i] =0 
	// 2.2.2 若 dp[i-1] > 0 && i - dp[i-1]>0 && s[i - dp[i-1] -1 ] == '(' 则 dp[i] = dp[i-1] + 2 + dp[i- dp[i-1] -2]
	// (subs) ( (sub_s) ) 
	// base case: dp[0] = 0
	// return max(dp[:])

	dp := make([]int, len(s))
	res := 0
	for i :=0; i<len(s); i++ {
		if i == 0 {
			dp[i] = 0
			continue
		}
		if i == 1 {
			if s[i] == ')' && s[i-1] == '(' {
				dp[i] = 2
				res = dp[i]
			}else {
				dp[i] = 0
			}
			continue
		}
		
		if s[i] == ')' && s[i-1] == '(' {
			dp[i] = dp[i-2] + 2
		}else if s[i] == ')' && s[i-1] == ')' &&
		 dp[i-1] > 0 && i - dp[i-1] -1 >= 0 && s[i - dp[i-1] -1 ] == '(' {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2>= 0 {
				dp[i] = dp[i] + dp[i- dp[i-1] -2]
			}
		}

		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
// @lc code=end

