/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */
package main

import "fmt"

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := []int{1}
	for i := 1; i < len(nums); i++ {
		j := i - 1
		lastMax := 0
		for ; j >= 0; j-- {
			if nums[j] < nums[i] {
				if lastMax < dp[j] {
					lastMax = dp[j]
				}
			}
		}
		dp = append(dp, lastMax+1)
	}
	max := 0
	for i := 0; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

// @lc code=end
func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
