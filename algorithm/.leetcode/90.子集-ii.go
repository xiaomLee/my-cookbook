/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode.cn/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (63.47%)
 * Likes:    1180
 * Dislikes: 0
 * Total Accepted:    340.4K
 * Total Submissions: 536.3K
 * Testcase Example:  '[1,2,2]'
 *
 * 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,2]
 * 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
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
 *
 *
 *
 *
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	var backtrack func(nums []int, pos int, track []int, res *[][]int)
	backtrack = func(nums []int, pos int, track []int, res *[][]int) {
		ans := make([]int, len(track))
		copy(ans, track)
		*res = append(*res, ans)

		for i := pos; i < len(nums); i++ {
			if i > pos && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			backtrack(nums, i+1, track, res)
			track = track[:len(track)-1]
		}
	}

	track := make([]int, 0)
	res := make([][]int, 0)
	sort.Ints(nums)
	backtrack(nums, 0, track, &res)
	return res
}

// @lc code=end

