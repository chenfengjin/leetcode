/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
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
//  key point：
// 1. you should remove the nth node from the end rather than the beginning
// 2. you should remove the next node, so add one node before node2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l2 := &ListNode{
		Next: head,
	}
	l1 := head

	cur1 := l1
	cur2 := l2
	for i := 0; i < n; i++ {
		cur1 = cur1.Next
	}
	for cur1 != nil {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	// l2 head guarantees that cur2.next is never nil
	cur2.Next = cur2.Next.Next

	return l2.Next

}

// @lc code=end
