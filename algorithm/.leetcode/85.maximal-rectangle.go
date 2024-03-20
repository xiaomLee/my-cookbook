/*
 * @lc app=leetcode.cn id=85 lang=golang
 * @lcpr version=30119
 *
 * [85] 最大矩形
 *
 * https://leetcode.cn/problems/maximal-rectangle/description/
 *
 * algorithms
 * Hard (55.05%)
 * Likes:    1631
 * Dislikes: 0
 * Total Accepted:    192.7K
 * Total Submissions: 350.1K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * 输出：6
 * 解释：最大矩形如上图所示。
 *
 *
 * 示例 2：
 *
 * 输入：matrix = [["0"]]
 * 输出：0
 *
 *
 * 示例 3：
 *
 * 输入：matrix = [["1"]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * rows == matrix.length
 * cols == matrix[0].length
 * 1 <= row, cols <= 200
 * matrix[i][j] 为 '0' 或 '1'
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	// h[i][j] 表示 以 matrix[i][j] 为右下角矩形顶点时，矩形的最大高度
	// 则 若 matrix[i][j] == '0' h[i][j] = 0 
	// matrix[i][j] == '1' h[i][j] = h[i-1][j] + 1

	// w[i][j] 表示以 matrix[i][j] 为右下角矩形顶点时，矩形的最大宽度
	// 则 若 matrix[i][j] == '0' w[i][j] = 0 
	// matrix[i][j] == '1' w[i][j] = min(w[i][:j]) && h[i][j] > 0

	// return max(w[i][j] * h[i][j])

	h := make([][]int, len(matrix))
	w := make([][]int, len(matrix))
	res := 0
	for i:=0; i<len(matrix); i++ {
		if h[i] == nil {
			h[i] = make([]int, len(matrix[0]))
		}
		if w[i] == nil {
			w[i] = make([]int, len(matrix[0]))
		}
		for j:=0; j<len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				h[i][j] = 0
				w[i][j] = 0
			}else {
				if i == 0 {
					h[i][j] = 0
				}else {
					h[i][j] = h[i-1][j] + 1
				}
				if j == 0 {
					w[i][j] = 0
				}else {
					temp := 0
					for k:=j; k>=0 && h[i][k] > 0; k-- {
						if 
					}
				}
			}

			// 计算area
			area := w[i][j] * h[i][j]

			fmt.Println(i,j, matrix[i][j]-'0', w[i][j], h[i][j], area)
			if area > res {
				res = area
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1"]]\n
// @lcpr case=end

*/

