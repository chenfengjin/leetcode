/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l3 := &ListNode{}
	cur := l3
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			cur.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		cur = cur.Next

	}
	if list2 != nil {
		list1, list2 = list2, list1
	}

	cur.Next = list1
	return l3.Next
}
