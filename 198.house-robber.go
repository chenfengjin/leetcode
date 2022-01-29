/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
// 重点在状态转移方程
// dp[i] = max(dp[i-2] + nums[i],dp[i-1])
// 还需要关注初始化条件
package main

import "fmt"

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	robbed := []bool{true, true}

	dp := []int{nums[0], nums[1]}
	// conner condition
	if nums[1] < nums[0] {
		dp = []int{nums[0], nums[0]}
		robbed = []bool{true, false}
	}

	for i := 2; i < len((nums)); i++ {
		if robbed[i-1] {
			if dp[i-1] < dp[i-2]+nums[i] {
				dp = append(dp, dp[i-2]+nums[i])
				robbed = append(robbed, true)
			} else {
				dp = append(dp, dp[i-1])
				robbed = append(robbed, false)
			}
		} else {
			dp = append(dp, dp[i-2]+nums[i])
			robbed = append(robbed, true)
		}
	}
	max := 0
	for i := 0; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	fmt.Println(dp)
	return max
}

// @lc code=end

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}
