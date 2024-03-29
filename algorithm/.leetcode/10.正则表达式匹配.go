/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 *
 * https://leetcode.cn/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (30.72%)
 * Likes:    3823
 * Dislikes: 0
 * Total Accepted:    403.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
 * 
 * 
 * '.' 匹配任意单个字符
 * '*' 匹配零个或多个前面的那一个元素
 * 
 * 
 * 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
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
 * 示例 2:
 * 
 * 
 * 输入：s = "aa", p = "a*"
 * 输出：true
 * 解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "ab", p = ".*"
 * 输出：true
 * 解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 20
 * 1 <= p.length <= 20
 * s 只包含从 a-z 的小写字母。
 * p 只包含从 a-z 的小写字母，以及字符 . 和 *。
 * 保证每次出现字符 * 时，前面都匹配到有效的字符
 * 
 * 
 */

// @lc code=start
func isMatch(s string, p string) bool {
	// dp[i][j] 表示 s[:i] p[:j] 能否匹配
	// 则存在如下状态转换
	// p[j-1] == s[i-1] 匹配掉当前字符 dp[i][j] = dp[i-1][j-1]
	// p[j-1] == '.' 仅能匹配一个任意字符 dp[i][j] = dp[i-1][j-1]
	// p[j-1] == '*' && p[j-2] != s[i-1] 匹配0个p[j-2]字符 dp[i][j] = dp[i][j-2]
	// p[j-1] == '*' && p[j-2] == s[i-1] 匹配0个p[j-2]字符 dp[i][j] = dp[i][j-2]
	// p[j-1] == '*' && p[j-2] == s[i-1] 匹配1个p[j-2]字符 dp[i][j] = dp[i-1][j-2]
	// p[j-1] == '*' && p[j-2] == s[i-1] 匹配多个p[j-2]字符 dp[i][j] = dp[i-1][j]
	// base case dp[0][j] = p[j-1]=='*' dp[i][0] = i==0

	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		if dp[i] == nil {
			dp[i] = make([]bool, len(p)+1)
		}
		for j := 0; j <= len(p); j++ {
			if j == 0 {
				dp[i][j] = i == 0
				continue
			}
			if i == 0 {
				dp[i][j] = false
				if j >= 2 && j%2 == 0 && p[j-1] == '*' {
					dp[i][j] = dp[i][j-2]
				}
				continue
			}
			if p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)][len(p)]

}
// @lc code=end

