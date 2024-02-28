/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode.cn/problems/jump-game/description/
 *
 * algorithms
 * Medium (43.35%)
 * Likes:    2633
 * Dislikes: 0
 * Total Accepted:    845.7K
 * Total Submissions: 2M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,3,1,1,4]
 * 输出：true
 * 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,2,1,0,4]
 * 输出：false
 * 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^4
 * 0 <= nums[i] <= 10^5
 * 
 * 
 */

// @lc code=start
func canJump(nums []int) bool {
	// dp[i] 表示能否到达 nums[i] 下标
	// dp[i] = range dp[:i-1] { return dp[j] && i-j<nums[j] }
	// base case dp[0] = true
	dp := make([]bool, len(nums))
	for i:=0; i<len(nums); i++ {
		if i == 0 {
			dp[0] = true
			continue
		}
		for j:=0; j < i; j++ {
			if dp[j] && i-j <= nums[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}
// @lc code=end

