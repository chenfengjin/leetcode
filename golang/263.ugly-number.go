/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] Ugly Number
 */
package main

// @lc code=start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 0 && n%5 == 0 {
		n /= 5
	}
	for n > 0 && n%2 == 0 {
		n /= 2
	}
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// @lc code=end
