/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
func minCut(s string) int {
	// dp[i] 表示 s[:i] 分割成回文串的最小分割次数
	// 则 dp[i] = min (i, dp[j]+1 && s[j:i] is huiwen)
	// base case dp[0] = 0 dp[1] = 0
	// return dp[i]
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 0, 0

	huiwen := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		if huiwen[i] == nil {
			huiwen[i] = make([]bool, len(s)+1)
		}
		for j := 0; j < i; j++ {
			huiwen[j][i] = isHuiWen(s[j:i])
		}
	}

	for i := 2; i <= len(s); i++ {
		if huiwen[0][i] {
			dp[i] = 0
			continue
		} else {
			dp[i] = i - 1
		}
		for j := 0; j < i; j++ {
			if huiwen[j][i] && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[len(s)]
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

