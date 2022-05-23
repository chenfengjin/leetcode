/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 100000
	collect := func(cur int) {
		if ret > cur {
			ret = cur
		}
	}
	cur := 1
	minDepthHelper(root, cur, collect)
	return ret
}
func minDepthHelper(root *TreeNode, cur int, collect func(int)) {
	if root.Left == nil && root.Right == nil {
		collect(cur)
		return
	}
	if root.Left != nil {
		minDepthHelper(root.Left, cur+1, collect)
	}
	if root.Right != nil {
		minDepthHelper(root.Right, cur+1, collect)
	}
}

// @lc code=end
