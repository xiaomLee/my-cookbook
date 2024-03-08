/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode.cn/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (43.19%)
 * Likes:    2187
 * Dislikes: 0
 * Total Accepted:    407.5K
 * Total Submissions: 943.5K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 * 
 * 测试用例的答案是一个 32-位 整数。
 * 
 * 子数组 是数组的连续子序列。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: nums = [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: nums = [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= nums.length <= 2 * 10^4
 * -10 <= nums[i] <= 10
 * nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
 * 
 * 
 */

// @lc code=start
func maxProduct(nums []int) int {
	// dp[i] 表示以 nums[i] 为结尾的最大连续子数组积
	// dp[i] = max(nums[i], dp[i-1]*nums[i])
	// return dp[lens(nums)-1]

	dp1 := make([]int, len(nums))
	dp2 := make([]int, len(nums))
	res := math.MinInt32
	for i:=0; i<len(nums); i++ {
		dp1[i] = nums[i]
		dp2[i] = nums[i]
		if i > 0 {
			dp1[i] = min(nums[i], dp1[i-1]*nums[i], dp2[i-1]*nums[i])
			dp2[i] = max(nums[i], dp1[i-1]*nums[i], dp2[i-1]*nums[i])
		}
		if dp2[i] > res {
			res = dp2[i]
		}
	}
	return res
}

func max(ints ...int) int {
	res := math.MinInt32
	for _, i := range ints {
		if i > res {
			res = i
		}
	}
	return res
}

func min(ints ...int) int {
	res := math.MaxInt32
	for _, i := range ints {
		if i < res {
			res = i
		}
	}
	return res
}

// @lc code=end

