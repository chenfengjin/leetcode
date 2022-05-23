/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] Fibonacci Number
 */
package main

// @lc code=start
func fib(n int) int {
	if n == 0 {
		return 0
	}
	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[len(dp)-1]
}

// @lc code=end
