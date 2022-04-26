/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change 2
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start
func change(amount int, coins []int) int {
	dp := []int{0}
	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && min > dp[i-coins[j]] {
				min = dp[i-coins[j]] + 1
			}
		}
		dp = append(dp, min)
	}
	if dp[len(dp)-1] == math.MaxInt32 {
		return -1
	}
	return dp[len(dp)-1]
}

// @lc code=end
func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
