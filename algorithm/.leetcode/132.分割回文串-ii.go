/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 *
 * https://leetcode.cn/problems/palindrome-partitioning-ii/description/
 *
 * algorithms
 * Hard (49.79%)
 * Likes:    730
 * Dislikes: 0
 * Total Accepted:    84.2K
 * Total Submissions: 169.2K
 * Testcase Example:  '"aab"'
 *
 * 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
 * 
 * 返回符合要求的 最少分割次数 。
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aab"
 * 输出：1
 * 解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "a"
 * 输出：0
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "ab"
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 仅由小写英文字母组成
 * 
 * 
 * 
 * 
 */

// @lc code=start
func minCut(s string) int {
	// dp[i] 表示将 s[:i] 分割成都是回文子串的最小分割次数 0<= i <= len(s)
	// 则 dp[i] = min(dp[j] && s[j:i] isPalindrome) + 1
	// base case dp[0] = 0 dp[1] = 0 dp[i] = i-1
	// return dp[len(s)]

	isPalindrome := func(s string) bool {
		for i,j:=0, len(s)-1; i<j; i, j= i+1, j-1{
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	huiwen := make([][]bool, len(s)+1)
	for i:=0; i<len(s); i++ {
		if huiwen[i] == nil {
			huiwen[i] = make([]bool, len(s)+1)
		}
		for j:=i; j<=len(s); j++ {
			huiwen[i][j] = isPalindrome(s[i:j])
		}
	}

	dp := make([]int, len(s)+1)
	for i:=0; i<=len(s); i++ {
		if i <= 1 || huiwen[0][i] {
			dp[i] = 0
			continue
		}

		dp[i] = i - 1
		for j:=0; j<i; j++ {
			if huiwen[j][i] && dp[j] + 1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[len(s)]
}
// @lc code=end

