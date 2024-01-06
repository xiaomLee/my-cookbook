/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode) {
	rebuild(root)
}

var pre *TreeNode

func rebuild(root *TreeNode) {
	if root == nil {
		return
	}

	rebuild(root.Left)
	if pre != nil {
		// tmp := pre
		// pre.Val = root.Val
		// root.Val = tmp.Val
		fmt.Println("pre:", pre.Val)
	}
	fmt.Println(root.Val)
	pre = root
	rebuild(root.Right)
}

// @lc code=end

