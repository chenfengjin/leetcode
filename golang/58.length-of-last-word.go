/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
package main

import "strings"

// @lc code=start
func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	array := strings.Split(s, " ")
	i := len(array) - 1
	for ; i >= 0; i-- {
		if len(array[i]) != 0 {
			break
		}
	}
	return len(array[i])
}

// @lc code=end
