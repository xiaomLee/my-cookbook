/*
 * @lc app=leetcode.cn id=82 lang=golang
 * @lcpr version=30119
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (54.19%)
 * Likes:    1279
 * Dislikes: 0
 * Total Accepted:    421.4K
 * Total Submissions: 777.4K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：head = [1,2,3,3,4,4,5]
 * 输出：[1,2,5]
 *
 *
 * 示例 2：
 *
 * 输入：head = [1,1,1,2,3]
 * 输出：[2,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围 [0, 300] 内
 * -100 <= Node.val <= 100
 * 题目数据保证链表已经按升序 排列
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	dummy.Next = head
	//  [null, 1,2,3,3,4,4,5] [1,2,5]
	// p == null p.Next=1 p.Next.Next=2
	p := dummy
	for p!=nil && p.Next!=nil && p.Next.Next!=nil {
		if p.Next.Val == p.Next.Next.Val {
			repeat := p.Next.Val
			for p.Next!=nil && p.Next.Val == repeat {
				p.Next = p.Next.Next
			}
		}else {
			p = p.Next
		}
		
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,3,4,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n
// @lcpr case=end

*/

