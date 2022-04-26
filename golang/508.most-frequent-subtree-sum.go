/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
 */
package main

//  it is not so fast

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
	values := map[int]int{}
	_, max := helper(root, &values, 0)
	ret := []int{}
	for key, value := range values {
		if value == max {
			ret = append(ret, key)
		}
	}
	return ret
}
func helper(root *TreeNode, value *map[int]int, max int) (int, int) {
	if root == nil {
		return 0, max
	}
	left, maxLeft := helper(root.Left, value, max)
	right, maxRight := helper(root.Right, value, max)
	sum := left + right + root.Val
	if _, ok := (*value)[sum]; ok {
		(*value)[sum] += 1
	} else {
		(*value)[sum] = 1
	}
	if max < (*value)[sum] {
		max = (*value)[sum]
	}
	if max < maxLeft {
		max = maxLeft
	}
	if max < maxRight {
		max = maxRight
	}
	return sum, max
}

// @lc code=end
