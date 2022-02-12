/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
//  it is a simulation to iterate over width
//  we wanna find the largest amount over with fixed width
package main

import "fmt"

// @lc code=start
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	cur := 0
	for left != right {
		if cur < min(height[left], height[right])*(right-left) {
			cur = min(height[left], height[right]) * (right - left)
		}
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}

	}
	return cur

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
func main() {
	fmt.Println(maxArea([]int{1, 3, 2, 5, 25, 24, 5}))
}
