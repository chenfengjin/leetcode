/*
 * @lc app=leetcode id=669 lang=golang
 *
 * [669] Trim a Binary Search Tree
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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	lowBound := root
	highBound := root
	for lowBound != nil && lowBound.Val > low {
		if lowBound.

	}

}

// @lc code=end
