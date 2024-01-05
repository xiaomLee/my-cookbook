/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 *
 * https://leetcode.cn/problems/combination-sum-iv/description/
 *
 * algorithms
 * Medium (52.73%)
 * Likes:    913
 * Dislikes: 0
 * Total Accepted:    167.7K
 * Total Submissions: 318K
 * Testcase Example:  '[1,2,3]\n4'
 *
 * 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
 *
 * 题目数据保证答案符合 32 位整数范围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3], target = 4
 * 输出：7
 * 解释：
 * 所有可能的组合为：
 * (1, 1, 1, 1)
 * (1, 1, 2)
 * (1, 2, 1)
 * (1, 3)
 * (2, 1, 1)
 * (2, 2)
 * (3, 1)
 * 请注意，顺序不同的序列被视作不同的组合。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [9], target = 3
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * nums 中的所有元素 互不相同
 * 1
 *
 *
 *
 *
 * 进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？
 *
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	return dpF(nums, target)

	res := 0
	backtrack(nums, target, 0, &res)
	return res
}

func backtrack(nums []int, target, trackSum int, res *int) {
	if trackSum == target {
		*res = *res + 1
		return
	}
	// nums 都为正数
	if trackSum > target {
		return
	}

	for i := 0; i < len(nums); i++ {
		trackSum += nums[i]
		backtrack(nums, target, trackSum, res)
		trackSum -= nums[i]
	}
}

func dpF(nums []int, target int) int {
	// dp[i] 表示 nums 所能凑出 target 组合数
	// 对每个 nums[i-1] 有两种选择：选取/不选取
	// dp[i] = dp[i-nums[i-1]] (选取) + dp[i] 不选取
	// base case： dp[0] = 1
	// return dp[target]

	dp := make([]int, target+1)

	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i == 0 {
				dp[i] = 1
				continue
			}
			dp[i] = dp[i]
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

// @lc code=end

