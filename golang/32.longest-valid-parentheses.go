/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */
package main

// @lc code=start
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	var dp []int
	if s[0:2] == "()" {
		dp = []int{0, 2}
	} else {
		dp = []int{0, 0}
	}
	idx := 2

	for idx < len(s) {
		last := dp[len(dp)-1]
		char := s[idx]
		if char == '(' {
			dp = append(dp, 0)
			idx += 1
			continue
		}
		if s[idx-1] == '(' {
			dp = append(dp, dp[len(dp)-2]+2)
			// key point 1
		} else if idx-last-1 >= 0 && s[idx-last-1] == '(' {
			cur := last + 2
			// key point 2
			if idx-last-2 >= 0 {
				cur += dp[idx-last-2]
			}
			dp = append(dp, cur)
		} else {
			dp = append(dp, 0)
		}
		idx += 1
	}

	max := dp[0]
	for _, item := range dp {
		if max < item {
			max = item
		}
	}
	return max
}

// @lc code=end
