/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode.cn/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (49.71%)
 * Likes:    2213
 * Dislikes: 0
 * Total Accepted:    764.8K
 * Total Submissions: 1.5M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]
 * 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出：[[1,6],[8,10],[15,18]]
 * 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,4],[4,5]]
 * 输出：[[1,5]]
 * 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	// 以 starti 为值对 intervals 进行排序
	mySort := func(intervals [][]int) {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
	}
	mySort(intervals)

	res := make([][]int, 0)
	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
			continue
		}
		pre := res[len(res)-1]
		if interval[0] > pre[1] {
			res = append(res, interval)
		} else if interval[1] > pre[1] {
			res[len(res)-1] = []int{pre[0], interval[1]}
		} else {
			// do nothing
		}
		// fmt.Println(pre, interval, res)
	}
	return res
}

// @lc code=end

