/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	layer := 0
	curLayerNode := []*TreeNode{root}
	result := [][]int{}
	for len(curLayerNode) != 0 {
		curLayerValue := []int{}
		nextLayerNode := []*TreeNode{}
		{ // for zigzag traveersal

			// zig zag means in the opposite direction everytime 
			len := len(curLayerNode)
			//  exchange nodes[i] and nodes[n-1-i] in place
			for i := 0; i < len/2; i++ {
				tmp := curLayerNode[i]
				curLayerNode[i] = curLayerNode[len-1-i]
				curLayerNode[len-1-i] = tmp
			}
		}
		for _, node := range curLayerNode {
			curLayerValue = append(curLayerValue, node.Val)
			if layer%2 == 0 {
				if node.Left != nil {
					nextLayerNode = append(nextLayerNode, node.Left)
				}
				if node.Right != nil {
					nextLayerNode = append(nextLayerNode, node.Right)
				}
			} else {
				if node.Right != nil {
					nextLayerNode = append(nextLayerNode, node.Right)
				}
				if node.Left != nil {
					nextLayerNode = append(nextLayerNode, node.Left)
				}

			}
		}
		result = append(result, curLayerValue)
		curLayerNode = nextLayerNode
		layer = layer + 1
	}
	return result
}

// @lc code=end

