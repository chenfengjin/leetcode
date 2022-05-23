/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] Counting Bits
 */
package main

// @lc code=start
func countBits(n int) []int {
	dp := []int{0}
	for len(dp) <= n {
		length := len(dp)
		for i := 0; i < length; i++ {
			dp = append(dp, dp[i]+1)
		}
	}
	return dp[:n+1]
}

// @lc code=end
