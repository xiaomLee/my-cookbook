/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode.cn/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (59.51%)
 * Likes:    1495
 * Dislikes: 0
 * Total Accepted:    483.9K
 * Total Submissions: 813.3K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用 一次 。
 *
 * 注意：解集不能包含重复的组合。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * 示例 2:
 *
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 *
 *
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	sum := 0
	sort.Ints(candidates)
	backtrack(candidates, target, 0, sum, track, &res)
	return res
}

func backtrack(nums []int, target, pos, sum int, track []int, res *[][]int) {
	if sum == target {
		ans := make([]int, len(track))
		copy(ans, track)
		*res = append(*res, ans)
		return // 1 <= candidates[i] <= 50
	}

	if sum > target {
		return
	}

	for i := pos; i < len(nums); i++ {
		if i > pos && nums[i] == nums[i-1] {
			continue
		}

		sum += nums[i]
		track = append(track, nums[i])

		backtrack(nums, target, i+1, sum, track, res)

		sum -= nums[i]
		track = track[:len(track)-1]
	}
}

// @lc code=end

