/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

package main

import (
	"fmt"
)

// @lc code=start
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	dp := []int{0}

	for i := 1; i < len(s); i++ {
		dp = append(dp, 0)
		char := string(s[i])
		if char == "(" {
			dp[i] = 0
			continue
		}

		if string(s[i-1]) == "(" {
			if i >= 2 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}

		} else {
			fmt.Println(i)
			if dp[i-1] < i && string(s[i-dp[i-1]-1]) == "(" {
				dp[i] = dp[i-1] + 2
				if i-1-dp[i-1]-1 > 0 {
					dp[i] += dp[i-1-dp[i-1]-1]
				}
			} else {
				dp[i] = 0
			}
		}

	}
	max := 0
	for _, i := range dp {
		if max < i {
			max = i
		}
	}
	return max
}

// @lc code=end
func main() {
	// fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses("(()())"))
}
