/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */
package main

// @lc code=start
func isPowerOfFour(n int) bool {
	mask := 0x55555554
	n &= mask
	return n != 0 && n&(n-1) == 0
}

// @lc code=end
