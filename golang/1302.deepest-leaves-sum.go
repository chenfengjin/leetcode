/*
 * @lc app=leetcode id=1302 lang=golang
 *
 * [1302] Deepest Leaves Sum
 */
package main

// COPY from 102
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := [][]int{}

	{
		curLayerNode := []*TreeNode{root}
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
	}

	layer := result[len(result)-1]
	ret := 0
	for _, node := range layer {
		ret += node
	}

	return ret
}

// @lc code=end
