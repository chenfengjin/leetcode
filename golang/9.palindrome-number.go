/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] Palindrome Number
 */
package main

// @lc code=start
func isPalindrome(x int) bool {
	origin := x
	palindrom := 0
	for x > 0 {
		digit := x % 10
		x /= 10
		palindrom *= 10
		palindrom += digit
	}
	return origin == palindrom
}

// @lc code=end
