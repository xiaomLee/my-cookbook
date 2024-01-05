/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// dp[i] 表示以 s[i] 为结尾能否被拆分
	// 则 dp[i] = for range s[:i] { dp[i] = dp[i] || (dp[j] && nums[j:i] in wordDict) }
	// base case dp[0] = wordDictMap[s[:i]]
	// return dp[len(s)-1]
	wordDictMap := make(map[string]struct{})
	for _, word := range wordDict {
		wordDictMap[word] = struct{}{}
	}
	fmt.Println(wordDictMap)
	dp := make([]bool, len(s)+1)

	for i := 0; i <= len(s); i++ {
		// fmt.Println(s[:i])
		if _, ok := wordDictMap[s[:i]]; ok {
			dp[i] = true
			continue
		}
		for j := 0; j < i; j++ {
			if _, ok := wordDictMap[s[j:i]]; ok && dp[j] {
				dp[i] = true
			}
		}
		// fmt.Println(dp[i])
	}
	return dp[len(s)]
}

// @lc code=end

