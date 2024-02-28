/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode.cn/problems/edit-distance/description/
 *
 * algorithms
 * Medium (62.79%)
 * Likes:    3266
 * Dislikes: 0
 * Total Accepted:    432.9K
 * Total Submissions: 689.3K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
 * 
 * 你可以对一个单词进行如下三种操作：
 * 
 * 
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= word1.length, word2.length <= 500
 * word1 和 word2 由小写英文字母组成
 * 
 * 
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
// @lc code=end

