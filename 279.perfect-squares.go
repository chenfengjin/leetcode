/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */
package main

import "fmt"

// @lc code=start
func numSquares(n int) int {

	table := make([]int, 101)

	for i := 0; i < 101; i++ {
		table[i] = i * i
	}
	right := 0
	dp := []int{1, 1, 2}
	for i := 3; i < n+1; i++ {
		for table[right] < i {
			right += 1
		}

		if i == table[right] {
			dp = append(dp, 1)
			continue
		}
		min := dp[i-1]
		for j := right; j > 0; j-- {
			if i-table[j] > 0 && min > dp[i-table[j]] {
				min = dp[i-table[j]]
			}
		}
		dp = append(dp, min+1)
	}
	return dp[n]
}

// @lc code=end

func main() {
	fmt.Println(numSquares(12))
}
