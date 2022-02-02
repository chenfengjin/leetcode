/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */
package main

import "fmt"

// TODO
// there is an mathematic method that runs faster
// @lc code=start
func integerBreak(n int) int {
	dp := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		max := i - 1
		for j := 1; j < i; j++ {
			if max < j*dp[i-j] {
				max = j * dp[i-j]
			}
			// key to solve this problem
			if max < j*(i-j) {
				max = j * (i - j)
			}

		}
		dp = append(dp, max)
	}
	fmt.Println(dp)
	return dp[len(dp)-1]
}

// @lc code=end

func main() {
	fmt.Print(integerBreak(4))
}
