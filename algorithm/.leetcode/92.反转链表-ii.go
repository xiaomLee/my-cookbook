/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode.cn/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (55.96%)
 * Likes:    1733
 * Dislikes: 0
 * Total Accepted:    473.5K
 * Total Submissions: 846.1K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left  。请你反转从位置 left 到位置 right 的链表节点，返回
 * 反转后的链表 。
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [5], left = 1, right = 1
 * 输出：[5]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点数目为 n
 * 1 
 * -500 
 * 1 
 * 
 * 
 * 
 * 
 * 进阶： 你可以使用一趟扫描完成反转吗？
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var tail *ListNode
func reverseN(head *ListNode, n int) *ListNode {
	// if n == 1 {
	// 	tail = head.Next
	// 	return head
	// }
	// newHead := reverseN(head.Next, n-1)
	// head.Next.Next = head
	// head.Next = tail
	// return newHead

	cur := head
	var pre *ListNode
	for cur!=nil && n>0 {
		cur.Next, pre, cur = pre, cur, cur.Next
		n--
	}
	head.Next = cur
	return pre
}
// @lc code=end

