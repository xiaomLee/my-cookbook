/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] 表示 text1[:i] text2[:j] 字串的最大公共子序列长度
	// 则 dp[i][j] = max(dp[i-1][j-1]+1 && text1[i] == text2[j], dp[i-1][j], dp[i][j-1])
	// base case dp[0][:], dp[:][0] = 0
	// return dp[len(text1)][len(text2)]

	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		if dp[i] == nil {
			dp[i] = make([]int, len(text2)+1)
		}
		for j := 0; j <= len(text2); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// @lc code=end

