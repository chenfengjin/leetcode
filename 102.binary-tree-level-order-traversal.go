/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
//  keep in sync with 104.maximum-depth-of-binary-tree
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	curLayerNode := []*TreeNode{root}
	result := [][]int{}
	for len(curLayerNode) != 0 {
		curLayerValue := []int{}
		nextLayerNode := []*TreeNode{}
		for _, node := range curLayerNode {
			curLayerValue = append(curLayerValue, node.Val)
			if node.Left != nil {
				nextLayerNode = append(nextLayerNode, node.Left)
			}
			if node.Right != nil {
				nextLayerNode = append(nextLayerNode, node.Right)
			}
		}
		result = append(result, curLayerValue)
		curLayerNode = nextLayerNode
	}
	return result
}

// @lc code=end
