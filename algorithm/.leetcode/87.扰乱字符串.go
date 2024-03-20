/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 *
 * https://leetcode.cn/problems/scramble-string/description/
 *
 * algorithms
 * Hard (47.12%)
 * Likes:    561
 * Dislikes: 0
 * Total Accepted:    62K
 * Total Submissions: 131.6K
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
 * 
 * 输入：s1 = "abcde", s2 = "caebd"
 * 输出：false
 * 
 * 
 * 示例 3：
 * 
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
 * 1 
 * s1 和 s2 由小写英文字母组成
 * 
 * 
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	// 区间dp：将 s1[i:i+k] s2[j:j+k] 当做一个整体区间
	// 每个字符串来说有两种选择：从第 i 位开始 截取 k 位长度作为一个虚拟整体
	// 则 能将 s1 s2 正确分割的条件为满足下述全部：
	// 1. s1[i:i+k] 与 s2[j:j+k] 能正确分割
	// 2. s1[:i] 与 (s2[:j] 或 s2[j+k:]) 且 s1[i+k:] 与(s2[:j] 或 s2[j+k:]) 能正确分割
	// 存在递归子问题
	// 故定义 dp[i][j][k] 为 分别从 s1, s2 的 i, j 位开始截取 k 长度的子串，此时的 s1[i:i+k] s2[j:k] 能否被正确分割
	// dp[i][j][k] = range(l ... k) { 
	// 		(dp[i][j][l] && dp[i+l][j+l][k-l]) ||	// (s1[i:i+l] && s2[j:j+l]) && (s1[i+l:i+k] && s2[j+l:j+k])
	//		(dp[i][j+k-l][l] && dp[i+l][j][k-l])	// (s1[i:i+l] && s2[j+k-l:j+k]) && (s1[i+l:i+k] && s2[j:j+k-l])
	// }
	// base case: 当 k==1 时 判断 s[i] == s[j] 即  dp[i][j][1] = s[i] == s[j]

	n := len(s1)
	dp := make([][][]bool, n+1)
	for k:=1; k <= n; k++ {
		for i:=0; i<=n-k; i++ {
			if dp[i] == nil {
				dp[i] = make([][]bool, n+1)
			}
			for j:=0; j<=n-k; j++ {
				if dp[i][j] == nil {
					dp[i][j] = make([]bool, n+1)
				}
				if k == 1 {
                    dp[i][j][k] = s1[i] == s2[j]
                    continue
                }
				for l:=1; l<k; l++ {
					// (s1[i:i+l] && s2[j:j+l]) && (s1[i+l:i+k] && s2[j+l:j+k])
					if dp[i][j][l] && dp[i+l][j+l][k-l] {
						dp[i][j][k] = true
						break
					}
					// (s1[i:i+l] && s2[j+k-l:j+k]) && (s1[i+l:i+k] && s2[j:j+k-l])
					if dp[i][j+k-l][l] && dp[i+l][j][k-l] {
						dp[i][j][k] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}
// @lc code=end

