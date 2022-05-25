/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */
package main

// @lc code=start
func trailingZeroes(n int) int {
	powersOfFive := []int{5, 25, 125, 125 * 5, 125 * 25, 125 * 125}
	ret := 0
	for _, item := range powersOfFive {
		ret += n / item
	}
	return ret
}

// @lc code=end
