/*
 * @lc app=leetcode id=382 lang=golang
 *
 * [382] Linked List Random Node
 */
package main

import "math/rand"

//  key is behaviour of rand function in standard library
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() int {
	ret := this.head.Val
	head := this.head
	count := 0
	for head != nil {
		count += 1
		if rand.Intn(count) == 0 {
			ret = head.Val
		}
		head = head.Next
	}
	return ret
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end
