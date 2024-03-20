/*
 * @lc app=leetcode.cn id=61 lang=golang
 * @lcpr version=30119
 *
 * [61] 旋转链表
 *
 * https://leetcode.cn/problems/rotate-list/description/
 *
 * algorithms
 * Medium (41.36%)
 * Likes:    1037
 * Dislikes: 0
 * Total Accepted:    366.5K
 * Total Submissions: 886.5K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
 *
 *
 *
 * 示例 1：
 *
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[4,5,1,2,3]
 *
 *
 * 示例 2：
 *
 * 输入：head = [0,1,2], k = 4
 * 输出：[2,0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 500] 内
 * -100 <= Node.val <= 100
 * 0 <= k <= 2 * 10^9
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
func rotateRight(head *ListNode, k int) *ListNode {
	// [1,2,3,4,5], k = 2 ==> [4,5,1,2,3]
	// 倒数第 k%n 个数翻转到开头即可
	var n = 0
	p := head
	for p!=nil {
		p = p.Next
		n++
	}
	if n == 0 || k%n == 0 {
		return head
	}
	slow, fast := head, head
	for k=k%n; k > 0; k-- {
		fast = fast.Next
	}
	// fmt.Println(fast.Val)
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// fmt.Println(slow.Val, fast.Val)
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2]\n4\n
// @lcpr case=end

*/

