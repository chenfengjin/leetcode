/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */
//  the key is coner case when n is less than or equal to zero
package main

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n < 0 {
		return false
	}
	for n&1 == 0 {
		n >>= 1
	}
	return n == 1 || n == -1
}

// @lc code=end
