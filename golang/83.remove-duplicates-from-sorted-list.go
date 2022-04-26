/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
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
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if cur == nil {
		return nil
	}
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// @lc code=end

// 1 2 3 4 5 6 6 7
// list
