/*
 * @lc app=leetcode.cn id=1161 lang=golang
 *
 * [1161] Maximum Level Sum of a Binary Tree
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
func sum(items []int) int {
	total := 0
	for i := 0; i < len(items); i++ {
		total += items[i]
	}
	return total

}
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
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
	max := -1 * 2 << 32
	ret := 0
	for idx, item := range result {
		if max < sum(item) {
			max = sum(item)
			ret = idx
		}
	}
	return ret + 1
}

// @lc code=end
