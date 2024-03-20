/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode.cn/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (73.81%)
 * Likes:    3501
 * Dislikes: 0
 * Total Accepted:    1.8M
 * Total Submissions: 2.4M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,2,3,4,5]
 * 输出：[5,4,3,2,1]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [1,2]
 * 输出：[2,1]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：head = []
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点的数目范围是 [0, 5000]
 * -5000 
 * 
 * 
 * 
 * 
 * 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
 * 
 * 
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
func reverseList(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// tail := reverseList(head.Next)
	// head.Next.Next = head
	// head.Next = nil
	// return tail
	
	var pre *ListNode
	for head != nil {
		// tmp := head
		// head = head.Next
		// tmp.Next = pre
		// pre = tmp

		head.Next, pre, head = pre, head, head.Next
		
	}
	return pre
}
// @lc code=end

