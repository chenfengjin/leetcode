/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */
package main

import (
	"fmt"
	"strings"
)

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
func binaryTreePaths(root *TreeNode) []string {
	ret := []string{}
	if root == nil {
		return ret
	}
	capture := func(path string) {
		ret = append(ret, path)
	}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}
	current := []string{}
	binaryTreePathHelper(root, capture, current)
	return ret
}
func binaryTreePathHelper(root *TreeNode, capture func(string), current []string) {
	if root.Left == nil && root.Right == nil {
		current = append(current, fmt.Sprintf("%d", root.Val))
		capture(strings.Join(current, "->"))
		return
	}
	current = append(current, fmt.Sprintf("%d", root.Val))
	if root.Left != nil {
		binaryTreePathHelper(root.Left, capture, current)
	}
	if root.Right != nil {
		binaryTreePathHelper(root.Right, capture, current)
	}
	current = current[:len(current)-1]
}

// @lc code=end
