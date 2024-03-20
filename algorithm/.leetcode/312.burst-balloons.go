/*
 * @lc app=leetcode.cn id=312 lang=golang
 * @lcpr version=30119
 *
 * [312] 戳气球
 *
 * https://leetcode.cn/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (69.99%)
 * Likes:    1322
 * Dislikes: 0
 * Total Accepted:    114.9K
 * Total Submissions: 164.1K
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
 * 输入：nums = [3,1,5,8]
 * 输出：167
 * 解释：
 * nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
 * coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
 *
 * 示例 2：
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

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxCoins(nums []int) int {
	// 套用回文串中心扩散的模型将本问题拆解成递归子问题
	// 将 nums[i:j] 当做一个整体被戳破能得到的最大硬币数是 dp[i][j]
	// 则 dp[i][j] = k range i...j max(dp[i][k] + dp[k][j] + nums[i]*nums[k]*nums[j])
	// base case dp[0][0] = nums[0]
	// return dp[0][len(nums)-1]

	// 为方便边界判断先给数组左右两边各加上1
	n := len(nums)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	// max := func(ints ...int) int {
	// 	sort.Ints(ints)
	// 	return ints[len(ints)-1]
	// }

	dp := make([][]int, n+2)
	for i := 0; i < n + 2; i++ {
        dp[i] = make([]int, n + 2)
    }

	for i:=n-1; i >=0; i-- {
		// i,j区间内至少有一个num
		// j 最大为新数组的最后一个数
		for j:=i+2; j<=n+1; j++ {
			for k:=i+1; k<j; k++ {
				sum := dp[i][k] + dp[k][j] + nums[i]*nums[k]*nums[j]
				if sum > dp[i][j] {
					dp[i][j] = sum
				}
				// dp[i][j] = max(dp[i][j], dp[i][k] + dp[k][j] + nums[i]*nums[k]*nums[j])
			}
		}
	}
	return dp[0][n+1]
}

// @lc code=end

/*
// @lcpr case=start
// [3,1,5,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,5]\n
// @lcpr case=end

*/

