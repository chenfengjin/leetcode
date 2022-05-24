/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] Jump Game
 */
package main

import "fmt"

// @lc code=start
func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		if !dp[i] {
			return false
		}
		for j := i; j <= i+nums[i]; j++ {
			if j < len(nums) {
				dp[j] = true
			}
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end
func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
