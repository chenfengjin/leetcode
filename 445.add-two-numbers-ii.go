/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
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
//  difference from, the mosg signification digit comes first
//  so we must calculate from end to start
//  a trivial solutions including
//  1. convert to int
//  2. convert to string
// stack is a good choice for reversing elements,we can use stack to store all elements
//  function call is a typical usage of stack
//  we can use recursion function call to solve this  problem

//  the trick of adding lead zero to the shorter list is not good
//  a user at leetcode-cn.com using linked list length to solve this problem
// see https://leetcode-cn.com/problems/add-two-numbers-ii/solution/java-2ms-ji-bai-100-by-she-hui-zhu-yi-jie-ban-ren-/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	length1 := 0
	length2 := 0

	list1 := l1
	list2 := l2

	for list1 != nil {
		length1++
		list1 = list1.Next
	}
	for list2 != nil {
		length2++
		list2 = list2.Next
	}
	// make sure  l1  is the longest
	if length1 < length2 {
		length2, length1 = length1, length2
		l1, l2 = l2, l1
	}
	for i := 0; i < length1-length2; i++ {
		node := ListNode{Val: 0}
		node.Next = l2
		l2 = &node
	}

	l3, carry := addOneNode(l1, l2)
	if carry == 1 {
		l3 = &ListNode{
			Val:  1,
			Next: l3,
		}
	}
	return l3
}
func addOneNode(l1 *ListNode, l2 *ListNode) (l3 *ListNode, carry int) {
	if l1 == nil || l2 == nil {
		return nil, 0
	}
	l3 = &ListNode{}

	l3.Next, carry = addOneNode(l1.Next, l2.Next)
	l3.Val = (l1.Val + l2.Val + carry) % 10
	carry = (l1.Val + l2.Val + carry) / 10
	return l3, carry
}

// @lc code=end
