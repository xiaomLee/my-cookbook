/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 *
 * https://leetcode.cn/problems/word-break-ii/description/
 *
 * algorithms
 * Hard (58.48%)
 * Likes:    741
 * Dislikes: 0
 * Total Accepted:    97.5K
 * Total Submissions: 166.8K
 * Testcase Example:  '"catsanddog"\n["cat","cats","and","sand","dog"]'
 *
 * 给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序
 * 返回所有这些可能的句子。
 * 
 * 注意：词典中的同一个单词可能在分段中被重复使用多次。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
 * 输出:["cats and dog","cat sand dog"]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入:s = "pineapplepenapple", wordDict =
 * ["apple","pen","applepen","pine","pineapple"]
 * 输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
 * 解释: 注意你可以重复使用字典中的单词。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
 * 输出:[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 
 * 
 * 1 <= s.length <= 20
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 10
 * s 和 wordDict[i] 仅有小写英文字母组成
 * wordDict 中所有字符串都 不同
 * 
 * 
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	track := make([]string, 0)
	res := make([]string, 0)
	backtrack(s, wordDict, 0, track, &res)
	return res
}

func backtrack(s string, wordDict []string, pos int, track []string, res *[]string) {
	if pos == len(s) {
		*res = append(*res, strings.Join(track, " "))
		return
	}
	if pos > len(s) {
		return
	}
	for _, word := range wordDict {
		if len(word) > len(s)-pos {
			continue
		}
		sub := s[pos:pos+len(word)]
		if sub != word {
			continue
		}
		track = append(track, sub)
		backtrack(s, wordDict, pos+len(word), track, res)
		track = track[:len(track)-1]
	}
}
// @lc code=end

