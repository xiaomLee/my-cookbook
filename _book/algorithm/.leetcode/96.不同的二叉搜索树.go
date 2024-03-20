/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	memo := make(map[string]int)
	return generate(1, n, memo)
}

func generate(lo, hi int, memo map[string]int) int {
	var ans int
	if lo > hi {
		ans = 1
		return ans
	}
	key := fmt.Sprintf("%s-%s", lo, hi)
	if count, ok := memo[key]; ok {
		return count
	}

	for i := lo; i <= hi; i++ {
		lefts := generate(lo, i-1, memo)
		rights := generate(i+1, hi, memo)
		ans += lefts * rights
	}
	memo[key] = ans
	return ans
}

// @lc code=end

