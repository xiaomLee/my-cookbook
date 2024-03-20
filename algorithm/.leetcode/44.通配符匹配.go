/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 *
 * https://leetcode.cn/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (33.93%)
 * Likes:    1125
 * Dislikes: 0
 * Total Accepted:    151.4K
 * Total Submissions: 446.2K
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给你一个输入字符串 (s) 和一个字符模式 (p) ，请你实现一个支持 '?' 和 '*' 匹配规则的通配符匹配：
 * 
 * 
 * '?' 可以匹配任何单个字符。
 * '*' 可以匹配任意字符序列（包括空字符序列）。
 * 
 * 
 * 
 * 
 * 判定匹配成功的充要条件是：字符模式必须能够 完全匹配 输入字符串（而不是部分匹配）。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aa", p = "a"
 * 输出：false
 * 解释："a" 无法匹配 "aa" 整个字符串。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "aa", p = "*"
 * 输出：true
 * 解释：'*' 可以匹配任意字符串。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "cb", p = "?a"
 * 输出：false
 * 解释：'?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= s.length, p.length <= 2000
 * s 仅由小写英文字母组成
 * p 仅由小写英文字母、'?' 或 '*' 组成
 * 
 * 
 */

// @lc code=start
func isMatch(s string, p string) bool {
	// dp[i][j] 表示 s[i-1] p[j-1] 是否匹配
	// 1. s[i-1] == p[j-1] dp[i][j] = dp[i-1][j-1]
	// 2. p[j-1] == '?' dp[i][j] = dp[i-1][j-1]
	// 3. p[j-1] == '*' dp[i][j] = dp[i-1][j] || dp[i-1][j-1] || dp[i][j-1]
	// 4. dp[i][j] = false
	// base case: dp[0][0] = true dp[0][1:] = dp[0][j-1] && p[j-1] == * dp[1:][0] = false
	// return dp[len(s)][len(p)]
	
	dp := make([][]bool, len(s)+1)
	for i:=0; i<=len(s); i++ {
		if dp[i] == nil {
			dp[i] = make([]bool, len(p)+1)
		}
		for j:=0; j<=len(p); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}
			if i ==0 && j>0 {
				if j == 1 {
					dp[0][j] = p[j-1] == '*'
				}else {
					dp[0][j] = dp[0][j-1] && p[j-1] == '*'
				}
				continue
			}
			if i>0 && j == 0 {
				dp[i][j] = false
				continue
			}

			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			}else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-1] || dp[i][j-1]
			}else {
				dp[i][j] = false
			}
		}
	}
	return dp[len(s)][len(p)]
}
// @lc code=end

