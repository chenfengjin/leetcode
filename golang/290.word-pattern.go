/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */
package main

import (
	"fmt"
	"strings"
)

// @lc code=start
func wordPattern(pattern string, s string) bool {
	pattern2s := map[string]string{}
	s2pattern := map[string]string{}

	sSplited := strings.Split(s, " ")
	patternSplited := strings.Split(pattern, "")
	if len(sSplited) != len(patternSplited) {
		return false
	}
	for i := 0; i < len(sSplited); i++ {
		if word, ok := pattern2s[patternSplited[i]]; !ok {
			pattern2s[patternSplited[i]] = sSplited[i]

		} else if word != sSplited[i] {
			return false
		}

		if char, ok := s2pattern[sSplited[i]]; !ok {
			s2pattern[sSplited[i]] = patternSplited[i]
		} else if char != patternSplited[i] {
			return false
		}
	}

	return true
}

// @lc code=end
func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
