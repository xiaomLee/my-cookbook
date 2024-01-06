/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode.cn/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (65.61%)
 * Likes:    1515
 * Dislikes: 0
 * Total Accepted:    513.4K
 * Total Submissions: 782.5K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 *
 *
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	var backtrack func(nums []int, visited []bool, track []int, res *[][]int)
	backtrack = func(nums []int, visited []bool, track []int, res *[][]int) {
		if len(track) == len(nums) {
			ans := make([]int, len(nums))
			copy(ans, track)
			*res = append(*res, ans)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] && !visited[i-1] {
				continue
			}
			track = append(track, nums[i])
			visited[i] = true
			backtrack(nums, visited, track, res)
			track = track[:len(track)-1]
			visited[i] = false
		}
	}

	visited := make([]bool, len(nums))
	track := make([]int, 0)
	res := make([][]int, 0)
	sort.Ints(nums)
	backtrack(nums, visited, track, &res)
	return res
}

// @lc code=end

