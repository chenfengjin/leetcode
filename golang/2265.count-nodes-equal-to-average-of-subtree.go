/*
 * @lc app=leetcode.cn id=2265 lang=golang
 *
 * [2265] Count Nodes Equal to Average of Subtree
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
func averageOfSubtree(root *TreeNode) int {
	ret := 0
	collect := func() {
		ret += 1
	}
	averageOfSubtreeHelper(root, collect)
	return ret
}
func averageOfSubtreeHelper(root *TreeNode, collect func()) (total, nodes int) {
	if root == nil {
		return 0, 0
	}
	leftTotal, leftNodes := averageOfSubtreeHelper(root.Left, collect)
	rightTotal, rightNodes := averageOfSubtreeHelper(root.Right, collect)
	total = leftTotal + rightTotal + root.Val
	nodes = leftNodes + rightNodes + 1

	if (total)/(nodes) == root.Val {
		collect()
	}
	return
}

// @lc code=end
