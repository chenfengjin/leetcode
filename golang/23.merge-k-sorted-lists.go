/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
package main

import "fmt"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func mergeKLists(lists []*ListNode) *ListNode {
// return &ListNode{}
// }

// @lc code=end

func main() {
	a := []int{1, 2, 3, 4}
	i := 2
	a = append(a[:i], a[i+1:]...)
	fmt.Print(a)
}
