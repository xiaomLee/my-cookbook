/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	// dp[i][j] 表示子串 word1[:i] word2[:j] 的最小编辑距离
	// 则 dp[i][j] = dp[i-1][j-1] && word1[i-1] == word2[j-1] || min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1 无操作，替换，删除，插入
	// base case dp[0][j] = j dp[i][0] = i
	// return dp[len(word1)][word2]

	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		if dp[i] == nil {
			dp[i] = make([]int, len(word2)+1)
		}
		for j := 0; j <= len(word2); j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(nums ...int) int {
	sort.Ints(nums)
	return nums[0]
}

// @lc code=end

