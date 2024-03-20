/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func generate(lo int, hi int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if lo > hi {
		res = append(res, nil)
		return res
	}

	// 以i为root 节点
	for i := lo; i <= hi; i++ {

		// lefts
		lefts := generate(lo, i-1)
		// rights
		rights := generate(i+1, hi)

		// 排列组合所有的lefts, rights
		for _, left := range lefts {
			for _, right := range rights {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return res
}

// @lc code=end

