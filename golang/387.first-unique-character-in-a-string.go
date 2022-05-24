/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */
package main

// @lc code=start
func firstUniqChar(s string) int {
	existMap := map[byte]int{}
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		existMap[char] += 1
	}
	for i := 0; i < len(s); i++ {
		if existMap[s[i]] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
