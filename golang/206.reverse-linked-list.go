/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead, tail := reverseListHelper(head)
	tail.Next = nil
	return newHead
}

func reverseListHelper(head *ListNode) (newHead, tail *ListNode) {
	if head.Next == nil {
		return head, head
	}
	newHead, tail = reverseListHelper(head.Next)
	tail.Next = head
	tail = tail.Next

	return newHead, tail

}

// @lc code=end
