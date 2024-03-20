/*
 * @lc app=leetcode.cn id=87 lang=golang
 * @lcpr version=30119
 *
 * [87] 扰乱字符串
 *
 * https://leetcode.cn/problems/scramble-string/description/
 *
 * algorithms
 * Hard (47.09%)
 * Likes:    561
 * Dislikes: 0
 * Total Accepted:    62.2K
 * Total Submissions: 132.2K
 * Testcase Example:  '"great"\n"rgeat"'
 *
 * 使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
 *
 * 如果字符串的长度为 1 ，算法停止
 * 如果字符串的长度 > 1 ，执行下述步骤：
 *
 * 在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y
 * 。
 * 随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y +
 * x 。
 * 在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。
 *
 *
 *
 *
 * 给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s1 = "great", s2 = "rgeat"
 * 输出：true
 * 解释：s1 上可能发生的一种情形是：
 * "great" --> "gr/eat" // 在一个随机下标处分割得到两个子字符串
 * "gr/eat" --> "gr/eat" // 随机决定：「保持这两个子字符串的顺序不变」
 * "gr/eat" --> "g/r / e/at" // 在子字符串上递归执行此算法。两个子字符串分别在随机下标处进行一轮分割
 * "g/r / e/at" --> "r/g / e/at" // 随机决定：第一组「交换两个子字符串」，第二组「保持这两个子字符串的顺序不变」
 * "r/g / e/at" --> "r/g / e/ a/t" // 继续递归执行此算法，将 "at" 分割得到 "a/t"
 * "r/g / e/ a/t" --> "r/g / e/ a/t" // 随机决定：「保持这两个子字符串的顺序不变」
 * 算法终止，结果字符串和 s2 相同，都是 "rgeat"
 * 这是一种能够扰乱 s1 得到 s2 的情形，可以认为 s2 是 s1 的扰乱字符串，返回 true
 *
 *
 * 示例 2：
 *
 * 输入：s1 = "abcde", s2 = "caebd"
 * 输出：false
 *
 *
 * 示例 3：
 *
 * 输入：s1 = "a", s2 = "a"
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * s1.length == s2.length
 * 1 <= s1.length <= 30
 * s1 和 s2 由小写英文字母组成
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

func isScramble(s1 string, s2 string) bool {
	return fun1(s1, s2)
}

func fun1(s1 string, s2 string) bool {
	// 递归 + 记忆搜索
	// tips:
	// 1. 可对 s1 s2 进行 hascode 进行初步判断筛选(长度、字母异位词)
	// 2. 对每次结果进行存储，防止重复计算
	memo := make(map[string]bool)
	var recursive func(s1, s2 string, memo map[string]bool) bool
	var hashCode func(s string) int 
	recursive = func (s1, s2 string, memo map[string]bool) bool {
		if len(s1) != len(s2) {
			return false
		}
		if hashCode(s1) != hashCode(s2) {
			return false
		}
		if s1==s2 {
			return true
		}
		if ans, ok := memo[s1+s2]; ok {
			return ans
		}
	
		n := len(s1)
		for i:=0; i<n-1; i++ {
			if recursive(s1[:i+1], s2[:i+1], memo) && recursive(s1[i+1:], s2[i+1:], memo) {
				memo[s1+s2] = true
				memo[s2+s1] = true
				return true
			}
			
			if recursive(s1[:i+1], s2[n-i-1:], memo) && recursive(s1[i+1:], s2[:n-i-1], memo) {
				memo[s1+s2] = true
				memo[s2+s1] = true
				return true
			}
		}
		memo[s1+s2] = false
		memo[s2+s1] = false
		return false
	}
	
	hashCode = func(s string) int {
		var code int
		for _, char := range s {
			code += 1<<(char - 'a')
		}
		return code
	}
	return recursive(s1, s2, memo)
}

func fun2(s1, s2 string) bool {
	// 区间动态规划
	// 
}


// @lc code=end

/*
// @lcpr case=start
// "great"\n"rgeat"\n
// @lcpr case=end

// @lcpr case=start
// "abcde"\n"caebd"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

*/

