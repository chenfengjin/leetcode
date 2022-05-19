/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := [][]int{}
	collect := func(item []int) {
		tmp := make([]int, len(item))
		copy(tmp, item)
		ret = append(ret, tmp)
	}
	cur := 0
	curPath := []int{}
	pathSumHelper(root, targetSum, cur, curPath, collect)

	return ret
}

func pathSumHelper(root *TreeNode, target int, cur int, currentPath []int, collect func([]int)) {
	if root == nil {
		return
	}

	currentPath = append(currentPath, root.Val)
	cur += root.Val

	if root.Left == nil && root.Right == nil {
		if cur == target {
			collect(currentPath)
		}
		return
	}

	pathSumHelper(root.Left, target, cur, currentPath, collect)
	pathSumHelper(root.Right, target, cur, currentPath, collect)

	currentPath = currentPath[:len(currentPath)-1]
}

// @lc code=end
