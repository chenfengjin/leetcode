/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] Ugly Number II
 */
package main

// @lc code=start
func nthUglyNumber(n int) int {
	dp := []int{1}

	idx2, idx3, idx5 := 0, 0, 0
	for len(dp) < n {
		cur := 0
		if dp[idx2]*2 < dp[idx3]*3 && dp[idx2]*2 < dp[idx5]*5 {
			cur = dp[idx2] * 2
			idx2++
		} else if dp[idx3]*3 < dp[idx5]*5 {
			cur = dp[idx3] * 3
			idx3++

		} else {
			cur = dp[idx5] * 5
			idx5++
		}
		if cur > dp[len(dp)-1] {
			dp = append(dp, cur)
		}
	}

	return dp[len(dp)-1]
}

// @lc code=end
