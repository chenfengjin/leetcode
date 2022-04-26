/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归内容
// 边界条件
// 终止条件
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}
	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 3}
	fmt.Println(isSameTree(p, q))
}

// @lc code=end
