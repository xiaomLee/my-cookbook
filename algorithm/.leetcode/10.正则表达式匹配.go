/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
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

