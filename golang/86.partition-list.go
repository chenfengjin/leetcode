/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		node := ListNode{Val: nums[i]}
		cur.Next = &node
		cur = cur.Next
	}
	return head
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	left := head

	for left.Next != nil && left.Next.Val < x {
		left = left.Next
	}
	if left.Next == nil {
		return head
	}
	right := left

	for right != nil && right.Next != nil {
		if right.Next.Val >= x {
			right = right.Next
		} else {
			node := right.Next
			right.Next = right.Next.Next

			node.Next = left.Next

			left.Next = node
			left = left.Next
		}

	}
	if right != nil && right.Val < x {
		// TODO
	}
	return head

}

// @lc code=end
func main() {
	l := NewLinkedList([]int{1, 4, 3, 2, 5, 2})
	l1 := partition(l, 3)
	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
}
