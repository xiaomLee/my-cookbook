/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 *
 * https://leetcode.cn/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (52.31%)
 * Likes:    1965
 * Dislikes: 0
 * Total Accepted:    485.1K
 * Total Submissions: 927.3K
 * Testcase Example:  '[1,5,11,5]'
 *
 * 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,5,11,5]
 * 输出：true
 * 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,5]
 * 输出：false
 * 解释：数组不能分割成两个元素和相等的子集。
 *
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
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	var backtrack func(nums []int, target, pos int, sum int) bool
	backtrack = func(nums []int, target, pos int, sum int) bool {
		if sum == target {
			return true
		}
		if sum > target {
			return false
		}
		for i := pos; i < len(nums); i++ {
			sum += nums[i]
			if backtrack(nums, target, i+1, sum) {
				return true
			}
			sum -= nums[i]
		}
		return false
	}
	// return backtrack(nums, target, 0, 0)

	// 定义 dp[i][j], 表示针对 nums[:i] 容量为 j, 能否恰好装满
	// 则 对于当前的 dp[i][j] 取决于当前数字 nums[i] 是否装入
	// 当前不装入，则从 nums[:i-1] 去选择： dp[i][j] = dp[i-1][j]
	// 当前装入，则 nums[:i-1] 可装入的容量变为 j- nums[i]： dp[i][j] = dp[i-1][j-nums[i]]
	// base case: dp[:][0] = true dp[0][1:] = false

	dp := make([][]bool, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		if dp[i] == nil {
			dp[i] = make([]bool, target+1)
		}
		for j := 0; j <= target; j++ {
			if j == 0 {
				dp[i][j] = true
				continue
			}
			if i == 0 {
				dp[i][j] = false
				continue
			}
			dp[i][j] = dp[i-1][j]
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][target]
}

// @lc code=end

