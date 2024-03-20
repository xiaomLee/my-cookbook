/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 *
 * https://leetcode.cn/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (69.99%)
 * Likes:    1318
 * Dislikes: 0
 * Total Accepted:    114.3K
 * Total Submissions: 163.3K
 * Testcase Example:  '[3,1,5,8]'
 *
 * 有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
 * 
 * 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i
 * - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
 * 
 * 求所能获得硬币的最大数量。
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,1,5,8]
 * 输出：167
 * 解释：
 * nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
 * coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,5]
 * 输出：10
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == nums.length
 * 1 <= n <= 300
 * 0 <= nums[i] <= 100
 * 
 * 
 */

// @lc code=start
func maxCoins(nums []int) int {
	// 区间dp
	// dp[i][j] 表示将 nums[i:j] 之间的气球全戳破时所能获取的最大值
	// 要获得最大值存在如下选择：对于 nums[i:j] 间的气球被戳破，会产生如下结果
	// dp[i][j]_case_k = dp[i][k] + nums[i]*nums[k]*nums[j] + dp[k][j]
	// 则 dp[i][j] = max(dp[i][j]_case_i+1...j-1)

	// 对原数组左右边界进行扩展 + 1，用于边界控制
	n := len(nums)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, n+2)
	for i:=0; i<len(dp); i++ {
		dp[i] = make([]int, n+2)
	}

	// 对于 dp[i][j]_case_n 的结果产生会依赖 case_i+k，故需倒序迭代
	// [1, 3,1,5,8, 1] 
	// max_j = n+2 == 1 
	// max_i = n == n-1+1 =  原数组的倒数第二个值
	for i:=n; i>=0; i-- {
		for j:=i+2; j<n+2; j++ {
			// 循环 nums[i:j]
			for k:=i+1; k<j; k++ {
				//fmt.Println("i, j, k:", i, j, k)
				case_k := dp[i][k] + nums[i]*nums[k]*nums[j] + dp[k][j]
				if case_k > dp[i][j] {
					dp[i][j] = case_k
				}
			}
		}
	}
	return dp[0][n+1]
}
// @lc code=end

