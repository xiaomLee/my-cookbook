/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 *
 * https://leetcode.cn/problems/word-break/description/
 *
 * algorithms
 * Medium (54.74%)
 * Likes:    2398
 * Dislikes: 0
 * Total Accepted:    521.5K
 * Total Submissions: 951.5K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
 * 
 * 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
 * 注意，你可以重复使用字典中的单词。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 300
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 20
 * s 和 wordDict[i] 仅由小写英文字母组成
 * wordDict 中的所有字符串 互不相同
 * 
 * 
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// dp[i] 表示 s[:i] 能否有单词组成
	// 则 dp[i] = j in {0, i} { dp[i] = dp[:j] && s[j:i] in Dict }
	// base case dp[0] = true
	// return dp[len(s)]

	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	dp := make([]bool, len(s)+1)
	for i:=0; i<=len(s); i++ {
		if i == 0 {
			dp[0] = true
			continue
		}
		for j:=0; j<i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
// @lc code=end

