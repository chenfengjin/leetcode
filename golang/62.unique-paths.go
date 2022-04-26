/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
package main

// @lc code=start
func uniquePaths(m int, n int) int {
	return C(m+n, m)
	// dp := make([][]int, m)
	// for i := 0; i < m; i++ {
	// 	dp[i] = make([]int, n)
	// }
	// for i := 0; i < m; i++ {
	// 	dp[i][0] = 1
	// }
	// for i := 0; i < n; i++ {
	// 	dp[0][i] = 1
	// }
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		// if i == 0 && j == 0 {
	// 		// 	dp[i][j] = 1
	// 		// } else if i == 0 {
	// 		// 	dp[i][j] = dp[i][j-1]
	// 		// } else if j == 0 {
	// 		// 	dp[i][j] = dp[i-1][j]
	// 		// } else {
	// 		dp[i][j] = dp[i-1][j] + dp[j-1][i]
	// 		// }
	// 	}
	// }
	// return dp[m-1][n-1]
}

// @lc code=end
