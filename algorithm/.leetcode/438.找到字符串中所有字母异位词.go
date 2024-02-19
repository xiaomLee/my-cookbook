/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (54.05%)
 * Likes:    1358
 * Dislikes: 0
 * Total Accepted:    361.1K
 * Total Submissions: 668.1K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
 * 
 * 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: s = "cbaebabacd", p = "abc"
 * 输出: [0,6]
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: s = "abab", p = "ab"
 * 输出: [0,1,2]
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= s.length, p.length <= 3 * 10^4
 * s 和 p 仅包含小写字母
 * 
 * 
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	win := make(map[byte]int, 0)
	res := make([]int, 0)
	left, right:=0, 0


	match := make(map[byte]int)
    for i:=0; i<len(p); i++ {
        match[p[i]] +=1
    }

	for right < len(s) {
		win[s[right]] += 1
		right++

		for right-left>len(p) {
            win[s[left]]--
            left++
        }

		if right-left == len(p) && valid(win, match) {
            res = append(res, left)
        }
	}
	return res
}

func valid(win map[byte]int, match map[byte]int) bool {
    for key, val := range match {
        if num, ok := win[key]; !ok || num != val {
            return false
        }
    }
    return true
}
// @lc code=end

