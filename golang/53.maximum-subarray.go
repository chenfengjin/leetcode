/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
package main

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp = append(dp, nums[i])
		} else {
			dp = append(dp, dp[i-1]+nums[i])
		}
	}

	max := dp[0]
	for i := 0; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

// @lc code=end
