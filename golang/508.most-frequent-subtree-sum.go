/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
 */
package main

import "fmt"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findFrequentTreeSum(root *TreeNode) []int {
	countMap := map[int]int{}
	collect := func(num int) {
		countMap[num] += 1
	}
	findFrequentTreeSumHelper(root, collect)
	ret := []int{}
	max := -1000000 * 100000
	for key, count := range countMap {
		if count > max {
			max = count
			ret = []int{key}

		} else if count == max {
			ret = append(ret, key)
		}
	}
	return ret
}
func findFrequentTreeSumHelper(root *TreeNode, collect func(int)) int {
	leftSum, rightSum := 0, 0
	if root.Left != nil {
		leftSum = findFrequentTreeSumHelper(root.Left, collect)
	}
	if root.Right != nil {
		rightSum = findFrequentTreeSumHelper(root.Right, collect)
	}
	collect(leftSum + rightSum + root.Val)
	return leftSum + rightSum + root.Val
}

// @lc code=end
