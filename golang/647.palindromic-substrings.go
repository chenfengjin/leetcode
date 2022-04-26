/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */
package main

import "fmt"

// @lc code=start
func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		left := i
		right := i
		// situation 1 s[i] is center of substring
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right++
		}

		// situation 2 s[i] is not center of substring, but end of left part of sub string

		left = i
		right = i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right++
		}
	}
	return count
}

// @lc code=end

func main() {
	fmt.Println(countSubstrings("aaa"))
}
