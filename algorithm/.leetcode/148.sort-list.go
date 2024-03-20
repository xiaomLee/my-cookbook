/*
 * @lc app=leetcode.cn id=148 lang=golang
 * @lcpr version=30119
 *
 * [148] 排序链表
 *
 * https://leetcode.cn/problems/sort-list/description/
 *
 * algorithms
 * Medium (65.52%)
 * Likes:    2271
 * Dislikes: 0
 * Total Accepted:    485.1K
 * Total Submissions: 740.2K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：head = [4,2,1,3]
 * 输出：[1,2,3,4]
 *
 *
 * 示例 2：
 *
 * 输入：head = [-1,5,3,4,0]
 * 输出：[-1,0,3,4,5]
 *
 *
 * 示例 3：
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
 * 链表中节点的数目在范围 [0, 5 * 10^4] 内
 * -10^5 <= Node.val <= 10^5
 *
 *
 *
 *
 * 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
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
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针找中点 归并排序
	// [4,2,1,3] 
	// 此处注意：fast = head.Next, 否则当只有两个数的时候，slow 指针会永远指向第二个数
	// slow 指针要寻找的相当于是 mid 的前一个节点
	slow, fast := head, head.Next
	for fast!= nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	fmt.Println(slow.Val, slow.Next.Val)
	right := sortList(slow.Next)
	slow.Next = nil
	left := sortList(head)
	return mergeTwoLists(left, right)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}else if list2 == nil {
		return list1
	}else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1,5,3,4,0]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

