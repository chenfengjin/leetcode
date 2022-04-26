/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */
// 1. 任何一种方案进行遍历，往下一层就加1
// 2. 层序遍历，看层数
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// keep in sync with 102.binary-tree-level-order-traversal.go
	curLayerNode := []*TreeNode{root}
	result := [][]int{}
	depth := 0
	for len(curLayerNode) != 0 {
		depth += 1
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
	return depth
}

// @lc code=end
