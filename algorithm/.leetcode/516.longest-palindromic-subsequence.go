/*
 * @lc app=leetcode.cn id=516 lang=golang
 * @lcpr version=30119
 *
 * [516] 最长回文子序列
 *
 * https://leetcode.cn/problems/longest-palindromic-subsequence/description/
 *
 * algorithms
 * Medium (67.10%)
 * Likes:    1178
 * Dislikes: 0
 * Total Accepted:    222.6K
 * Total Submissions: 331.8K
 * Testcase Example:  '"bbbab"'
 *
 * 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
 *
 * 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "bbbab"
 * 输出：4
 * 解释：一个可能的最长回文子序列为 "bbbb" 。
 *
 *
 * 示例 2：
 *
 * 输入：s = "cbbd"
 * 输出：2
 * 解释：一个可能的最长回文子序列为 "bb" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由小写英文字母组成
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestPalindromeSubseq(s string) int {
	// dp[i][j] 表示 s[j:i] 的最长回文子序列
	// 1. s[i] == s[j] dp[i][j] = max(dp[i-1][j+1]+2, dp[i-1][j], dp[i][j+1])
	// 2. s[i] != s[j] dp[i][j] = max(dp[i-1][j+1], dp[i-1][j], dp[i][j+1])
	// base case: dp[0][0] = 1 
	// return dp[0][len(s)-1]
	
	max := func(ints ...int) int {
		sort.Ints(ints)
		return ints[len(ints)-1]
	}

	dp := make([][]int, len(s))
	for i:=0; i<len(s); i++ {
		if dp[i] == nil {
			dp[i] = make([]int, len(s))
		}
		for j:=i; j>=0; j-- {
			if i==0 || j==i {
				dp[i][j] = 1
				continue
			}
			if s[j] == s[i] {
				dp[i][j] = max(dp[i-1][j+1]+2, dp[i-1][j], dp[i][j+1])
			}else {
				dp[i][j] = max(dp[i-1][j+1], dp[i-1][j], dp[i][j+1])
			}
		}
	}
	return dp[len(s)-1][0]
}

// @lc code=end

/*
// @lcpr case=start
// "bbbab"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

