/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] Path Sum
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, target int) bool {
	return hasPathSumHelper(root, target, 0)
}

func hasPathSumHelper(root *TreeNode, target int, cur int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if cur+root.Val == target {
			return true
		}
		return false
	}
	left := hasPathSumHelper(root.Left, target, cur+root.Val)
	right := hasPathSumHelper(root.Right, target, cur+root.Val)
	return left || right
}

// @lc code=end
