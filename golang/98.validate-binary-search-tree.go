/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	items := []int{}
	ret := true

	collect := func(i int) {
		if len(items) > 0 && items[len(items)-1] >= i {
			ret = false
		}
		items = append(items, i)
	}
	isValidBSTHelper(root, collect)
	return ret
}

func isValidBSTHelper(root *TreeNode, collect func(int)) {
	if root.Left == nil && root.Right == nil {
		collect(root.Val)
		return
	}
	if root.Left != nil {
		isValidBSTHelper(root.Left, collect)
	}

	collect(root.Val)

	if root.Right != nil {
		isValidBSTHelper(root.Right, collect)
	}

}

// @lc code=end
