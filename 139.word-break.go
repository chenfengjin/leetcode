/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */
package main

import "fmt"

// @lc code=start
// to be honest, it is not a typic dynamic programming
// a important key to solve this problem is contraints 3 that wordDict[i].lenght <=20
func wordBreak(s string, wordDict []string) bool {
	dp := []bool{true}
	for i := 0; i < len(s); i++ {
		matched := false
		for j := 0; j < len(wordDict); j++ {
			word := wordDict[j]
			length := len(wordDict[j])
			if i-length+1 >= 0 && i-length < len(dp) && dp[i-length+1] {
				if s[i-length+1:i+1] == word {
					matched = true
					break
				}
			}
		}
		dp = append(dp, matched)
	}
	fmt.Println(dp)
	return dp[len(s)]
}

// @lc code=end

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
