/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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
func maxPathSum(root *TreeNode) int {
	max, maxSubtreeSum := helper(root)
	if max < maxSubtreeSum {
		return maxSubtreeSum
	}
	return max
}
func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, -9999
	}
	maxLeft, maxPathLeft := helper(root.Left)
	maxRight, maxPathRight := helper(root.Right)

	maxSubtreeSum := maxPathLeft
	if maxSubtreeSum < maxPathRight {
		maxSubtreeSum = maxPathRight
	}

	max := maxPathLeft + maxPathRight + root.Val
	if max < maxLeft {
		max = maxLeft
	}
	if max < maxRight {
		max = maxRight
	}

	return max, maxSubtreeSum + root.Val
}

// @lc code=end
