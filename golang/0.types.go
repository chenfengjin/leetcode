package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

// func maxIntSlice([]int) {

// }

// func NewLinkedList(nums []int) *ListNode {
// 	head := &ListNode{Val: nums[0]}
// 	cur := head
// 	for i := 1; i < len(nums); i++ {
// 		node := ListNode{Val: nums[i]}
// 		cur.Next = &node
// 		cur = cur.Next
// 	}
// 	return head
// }
