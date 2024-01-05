/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
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

