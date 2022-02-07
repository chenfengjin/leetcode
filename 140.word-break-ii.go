/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */
package main

import (
	"fmt"
)

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	dp := []bool{true}
	segs_dp := map[int][]string{
		// 0: []string{},
	}
	for i := 0; i < len(s); i++ {
		matched := false
		segs := []string{}

		for j := 0; j < len(wordDict); j++ {
			word := wordDict[j]
			length := len(wordDict[j])
			if i-length+1 >= 0 && i-length < len(dp) && dp[i-length+1] {
				if s[i-length+1:i+1] == word {
					matched = true
					if i-length+1 == 0 {
						segs = append(segs, word)

					} else {
						for _, seg := range segs_dp[i-length] {
							segs = append(segs, seg+" "+word)
						}
					}

				}
			}
		}
		dp = append(dp, matched)
		segs_dp[i] = segs
	}
	fmt.Println(segs_dp)
	return segs_dp[len(s)-1]
}

// @lc code=end
func main() {
	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))

}
