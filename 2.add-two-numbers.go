/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	head := sum
	carry := 0
	for l1 != nil && l2 != nil {

		val := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		l1 = l1.Next
		l2 = l2.Next

		node := ListNode{Val: val}
		head.Next = &node

		head = head.Next
	}
	if l1 != nil {
		l2 = l1
	}

	for l2 != nil {
		val := (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10
		head.Next = &ListNode{Val: val}
		head = head.Next
		l2 = l2.Next
	}
	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}
	return sum.Next
}

// @lc code=end
