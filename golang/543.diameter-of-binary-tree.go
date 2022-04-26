/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */
package main

// keys
// 1. pay attention go global variables usage
// 2. pay attention to defination of distance

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	maxDepth543 = 0
)

func diameterOfBinaryTree(root *TreeNode) int {
	maxDepth543 = 0
	helper543(root)
	return maxDepth543
}
func helper543(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	max := 0

	left := helper543(root.Left)
	right := helper543(root.Right)

	if root.Left == nil {
		max = right + 1
	} else if root.Right == nil {
		max = left + 1
	} else {
		max = left + right + 2
	}

	if maxDepth543 < max {
		maxDepth543 = max
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// @lc code=end
