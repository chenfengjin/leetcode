/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	carry := 0
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		val := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		cur.Next = &ListNode{Val: val}

		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		l1 = l2
	}

	for l1 != nil {
		val := (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		cur.Next = &ListNode{Val: val}

		cur = cur.Next
		l1 = l1.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return head.Next
}

// @lc code=end

