/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
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
func insertionSortList(head *ListNode) *ListNode {
	ret := &ListNode{}
	for head != nil {
		cur := ret
		for cur.Next.Val < head.Val {
			cur = cur.Next
		}
		tmp := cur.Next
		cur.Next

	}

}

// @lc code=end
