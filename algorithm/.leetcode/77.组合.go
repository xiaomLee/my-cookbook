/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode.cn/problems/combinations/description/
 *
 * algorithms
 * Medium (77.05%)
 * Likes:    1566
 * Dislikes: 0
 * Total Accepted:    636.6K
 * Total Submissions: 826.2K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
 *
 * 你可以按 任何顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4, k = 2
 * 输出：
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 * 示例 2：
 *
 *
 * 输入：n = 1, k = 1
 * 输出：[[1]]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */

// @lc code=start
func combine(n int, k int) [][]int {
	if k > n {
		return nil
	}

	var backtrack func(n, k int, pos int, track []int, res *[][]int)
	backtrack = func(n, k int, pos int, track []int, res *[][]int) {
		if len(track) == k {
			ans := make([]int, k)
			copy(ans, track)
			*res = append(*res, ans)
		}
		for i := pos; i <= n; i++ {
			track = append(track, i)
			backtrack(n, k, i+1, track, res)
			track = track[:len(track)-1]
		}
	}

	track := make([]int, 0)
	res := make([][]int, 0)
	backtrack(n, k, 1, track, &res)
	return res
}

// @lc code=end

