/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] Length of Last Word
 */
package main

import "strings"

// @lc code=start
func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	idx := len(words) - 1
	for len(words[idx]) == 0 {
		idx--
	}
	return len(words[idx])
}

// @lc code=end
