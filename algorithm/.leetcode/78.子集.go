/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode.cn/problems/subsets/description/
 *
 * algorithms
 * Medium (81.16%)
 * Likes:    2220
 * Dislikes: 0
 * Total Accepted:    714K
 * Total Submissions: 879.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 * nums 中的所有元素 互不相同
 *
 *
 */

// @lc code=start
func subsets(nums []int) [][]int {
	var backtrack func(nums []int, pos int, track []int, res *[][]int)
	backtrack = func(nums []int, pos int, track []int, res *[][]int) {
		ans := make([]int, len(track))
		copy(ans, track)
		*res = append(*res, ans)

		for i := pos; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(nums, i+1, track, res)
			track = track[:len(track)-1]
		}
	}

	track := make([]int, 0)
	res := make([][]int, 0)
	backtrack(nums, 0, track, &res)
	return res
}

// @lc code=end

