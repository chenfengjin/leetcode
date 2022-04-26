/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
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
func reverseList(head *ListNode) *ListNode {
	return head
}

// @lc code=end

func newLinkedList(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val: nums[i]}
		cur.Next = node
		cur = cur.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func traversal(list *ListNode) []int {
	return []int{}
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	head := newLinkedList(nums)
	newHead := reverseList(head)
	newNums := traversal(newHead)
	fmt.Println(newNums)
}
