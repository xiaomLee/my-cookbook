/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode.cn/problems/permutations/description/
 *
 * algorithms
 * Medium (78.98%)
 * Likes:    2785
 * Dislikes: 0
 * Total Accepted:    981.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[[0,1],[1,0]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * nums 中的所有整数 互不相同
 *
 *
 */

// @lc code=start
func permute(nums []int) [][]int {
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
	backtrack(nums, visited, track, &res)
	return res
}

// @lc code=end

