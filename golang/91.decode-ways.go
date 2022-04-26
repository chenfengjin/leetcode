/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
package main

import (
	"fmt"
	"strconv"
)

//TODO  try to make code clean
//  look ahead may be a good idea
//  key is how to deal with zero
//
// @lc code=start
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	dp := []int{1}
	n, _ := strconv.Atoi(s[0:2])

	if n%10 == 0 && n != 20 && n != 10 {
		return 0
	}
	if n > 26 || n == 10 || n == 20 {
		dp = append(dp, 1)
	} else {
		dp = append(dp, 2)
	}

	for i := 2; i < len(s); i++ {
		//  it may be an illegal string

		//  handle zero
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			} else {
				dp[i-1] = dp[i-2]
				dp = append(dp, dp[i-1])
				continue
			}
		}
		n, _ := strconv.Atoi(s[i-1 : i+1])
		if n > 26 || s[i-1] == '0' {
			dp = append(dp, dp[i-1])
		} else {
			dp = append(dp, dp[i-1]+dp[i-2])
		}
	}
	fmt.Println(dp)
	return dp[len(s)-1]
}

// @lc code=end
func main() {
	fmt.Println(numDecodings("226"))
}
