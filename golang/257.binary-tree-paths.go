/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */
package main

import (
	"fmt"
	"strings"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	ret := []string{}
	collect := func(items []int) {
		itemStr := []string{}
		for _, item := range items {
			itemStr = append(itemStr, fmt.Sprintf("%d", item))
		}
		ret = append(ret, strings.Join(itemStr, "->"))
	}

	cur := []int{}

	binaryTreePathsHelper(root, cur, collect)
	return ret

}
func binaryTreePathsHelper(root *TreeNode, cur []int, collect func([]int)) {
	cur = append(cur, root.Val)

	if root.Left == nil && root.Right == nil {
		collect(cur)
		return
	}

	if root.Left != nil {
		binaryTreePathsHelper(root.Left, cur, collect)
	}
	if root.Right != nil {
		binaryTreePathsHelper(root.Right, cur, collect)
	}
}

// @lc code=end
