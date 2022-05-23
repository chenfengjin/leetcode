/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] Power of Three
 */
package main

// @lc code=start
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	max := 1
	for max*3 < 1<<32 {
		max *= 3
	}
	return max%n == 0
}

// @lc code=end
