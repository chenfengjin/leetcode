/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] Path Sum III
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
//  官方解答更暴力
func pathSum(root *TreeNode, targetSum int) int {
	ret := 0
	cur := []int{}
	collect := func(items []int) {
		total := 0
		for i := len(items) - 1; i >= 0; i-- {
			total += items[i]
			if total == targetSum {
				ret++
			}
		}
	}
	pathSumHelper2(root, cur, collect)
	return ret
}

func pathSumHelper2(root *TreeNode, cur []int, collect func([]int)) {
	if root == nil {
		return
	}
	cur = append(cur, root.Val)
	collect(cur)

	if root.Left != nil {
		pathSumHelper2(root.Left, cur, collect)
	}
	if root.Right != nil {
		pathSumHelper2(root.Right, cur, collect)
	}
	// 不必要回溯
}

// @lc code=end
