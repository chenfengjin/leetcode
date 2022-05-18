/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	collect := func(delta int) {
		ret += delta
	}
	cur := 0
	sumNumberHelper(root, cur, collect)
	return ret

}

func sumNumberHelper(root *TreeNode, cur int, collect func(int)) {
	if root.Left == nil && root.Right == nil {
		collect(cur*10 + root.Val)
		return
	}
	cur = cur*10 + root.Val
	if root.Left != nil {
		sumNumberHelper(root.Left, cur, collect)
	}
	if root.Right != nil {
		sumNumberHelper(root.Right, cur, collect)
	}

}

// @lc code=end
