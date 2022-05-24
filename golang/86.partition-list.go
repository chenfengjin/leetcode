/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] Partition List
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
//  there is a solution that is so easy with only cost of tow ListNode
//  tow pointer is too complex and not neccessary
func partition(head *ListNode, x int) *ListNode {
	headLess := &ListNode{}
	headGreater := &ListNode{}

	curLess := headLess
	curGreater := headGreater

	cur := head
	for cur != nil {
		if cur.Val < x {
			curLess.Next = cur
			curLess = curLess.Next
		} else {
			curGreater.Next = cur
			curGreater = curGreater.Next
		}
		cur = cur.Next
	}
	curLess.Next = headGreater.Next
	curGreater.Next = nil
	return headLess.Next
}

// @lc code=end
