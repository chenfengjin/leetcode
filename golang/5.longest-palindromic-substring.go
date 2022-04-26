/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
package main

import "fmt"

// @lc code=start
func longestPalindrome(s string) string {
	maxLength := 0
	longestSubstr := ""
	for i := 0; i < len(s); i++ {
		left := i
		right := i
		// situation 1 s[i] is center of substring
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if maxLength < right-left+1 {
				maxLength = right - left + 1
				longestSubstr = s[left : right+1]
			}
			left--
			right++
		}

		// situation 2 s[i] is not center of substring, but end of left part of sub string

		left = i
		right = i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if maxLength < right-left+1 {
				maxLength = right - left + 1
				longestSubstr = s[left : right+1]
			}
			left--
			right++
		}

	}
	return longestSubstr
}

// @lc code=end

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("112211"))
	fmt.Println(longestPalindrome("cbbd"))

}
