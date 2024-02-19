/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode.cn/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (44.79%)
 * Likes:    983
 * Dislikes: 0
 * Total Accepted:    279.9K
 * Total Submissions: 624.8K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
 * 
 * 换句话说，s1 的排列之一是 s2 的 子串 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s1 = "ab" s2 = "eidbaooo"
 * 输出：true
 * 解释：s2 包含 s1 的排列之一 ("ba").
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s1= "ab" s2 = "eidboaoo"
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s1.length, s2.length <= 10^4
 * s1 和 s2 仅包含小写字母
 * 
 * 
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	s1Map := make(map[byte]int)
	for i:=0; i<len(s1); i++ {
		s1Map[s1[i]] += 1
	}

	match := func(m1, m2 map[byte]int) bool {
		for c, val := range m2 {
			if num, ok := m1[c]; !ok || num != val {
				return false
			}
		}
		return true
	}

	win := make(map[byte]int)
	left, right := 0, 0

	for right < len(s2) {
		win[s2[right]] += 1
		right++

		for right - left > len(s1) {
			win[s2[left]]--
			left++
		}

		// fmt.Println(win)

		if right - left == len(s1) && match(win, s1Map) {
			return true
		}
	}
	return false
}
// @lc code=end

