/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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
func generateTrees(n int) []*TreeNode {
	return helper95(1, n)

}

func helper95(m, n int) []*TreeNode {
	if m >= n {
		return []*TreeNode{
			nil,
		}
	}
	ret := []*TreeNode{}
	for i := m; i < n; i++ {
		root := TreeNode{Val: i}
		lefts := helper95(m, i)
		rights := helper95(i+1, n)
		for _, left := range lefts {
			for _, right := range rights {
				root.Left = left
				root.Right = right
				ret = append(ret, &root)
			}
		}
	}
	return ret
}

// @lc code=end
