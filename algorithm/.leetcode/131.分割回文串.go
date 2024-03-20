/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode.cn/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (73.47%)
 * Likes:    1715
 * Dislikes: 0
 * Total Accepted:    352.3K
 * Total Submissions: 479.5K
 * Testcase Example:  '"aab"'
 *
 * 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
 * 
 * 回文串 是正着读和反着读都一样的字符串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aab"
 * 输出：[["a","a","b"],["aa","b"]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "a"
 * 输出：[["a"]]
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
 */

// @lc code=start
func partition(s string) [][]string {
	var backtrack func(s string, pos int, track []string, res *[][]string) 
	backtrack = func(s string, pos int, track []string, res *[][]string) {
		if pos == len(s) {
			ans := make([]string, len(track))
			copy(ans, track)
			*res = append(*res, ans)
			return
		}
		for i:=pos; i<len(s); i++ {
			// 对于每一个 s[i] 都有两种选择：分割 or 不分割
			if !isPalindrome(s[pos:i+1]) {
				continue
			}
			track = append(track, s[pos:i+1])
			backtrack(s, i+1, track, res)
			track = track[:len(track)-1]
		}
	}

	track := make([]string, 0)
	res := make([][]string, 0)
	backtrack(s, 0, track, &res)
	return res
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	i, j:= 0, len(s)-1
	for i<j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end

