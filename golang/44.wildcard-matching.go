/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */
//  dp[i][j] means p[0:i] match  s[0:j]

package main

import (
	"fmt"
)

//  注意 * 可以匹配 0 个或者多个字符
// @lc code=start
func isMatch(s string, p string) bool {
	if s == "" {
		for _, c := range p {
			if c != '*' {
				return false
			}
		}
		return p==""

	}

	if p == "" {
		return false
	}

	dp := [][]bool{}

	for i := 0; i < len(p); i++ {
		line := make([]bool, len(s))
		dp = append(dp, line)
	}

	dp[0][0] = s[0] == p[0] || p[0] == '?' || p[0] == '*'

	if p[0] == '*' {
		for i := 1; i < len(s); i++ {
			dp[0][i] = true
		}
	}

	for j := 1; j < len(p); j++ {
		if !dp[j-1][0] {
			continue
		}

		flag := true
		if p[j] == '*' {
			for l := 0; l < j; l++ {
				if p[l] != '*' {
					flag = false
					break
				}
			}
			dp[j][0] = flag

		} else if p[j] == '?' {
			for l := 0; l < j; l++ {
				if p[l] != '*' {
					flag = false
					// dp[0][j] = false
					break
				}
			}
			dp[j][0] = flag
		} else if p[j] == s[0] {
			for l := 0; l < j; l++ {
				if p[l] != '*' {
					flag = false
					// dp[0][j] = false
					break
				}
			}
			dp[j][0] = flag
		}
	}
	fmt.Println(dp[0])
	for i := 1; i < len(p); i++ {
		for j := 1; j < len(s); j++ {
			if p[i] == '*' {
				dp[i][j] = dp[i-1][j-1] || dp[i-1][j]
			} else if p[i] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[i] == s[j] {
				dp[i][j] = dp[i-1][j-1]

			}
		}
		fmt.Println(dp[i])
	}
	return dp[len(p)-1][len(s)-1]
}

// @lc code=end
func main() {

	// fmt.Println(isMatch("abcabczzzde","*abc???de*"))
	// fmt.Println(isMatch("","****"))
	fmt.Println(isMatch("mississippi", "m??*ss*?i*pi"))

}
