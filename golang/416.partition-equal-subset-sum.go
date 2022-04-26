/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */
package main

import "fmt"

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	sum >>= 1
	total := 0
	targets := make([]int, len(nums))
	targets[(len(nums) - 1)] = 0
	for i := len(targets) - 2; i >= 0; i-- {
		targets[i] = targets[i+1] + nums[i+1]
	}
	dp := make([]bool, len(nums))
	dp[0] = targets[0] == sum
	for i := 1; i < len(nums); i++ {

	}
}

func canPartitionRecursionWithTLE(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	sum >>= 1
	// total := 0
	// targets := make([]int, len(nums))
	// targets[(len(nums) - 1)] = 0
	// for i := len(targets) - 2; i >= 0; i-- {
	// 	targets[i] = targets[i+1] - nums[i+1]
	// }
	return helper416(nums, sum)
}

//  helper416 determine nums can sum to an target
func helper416(nums []int, target int) bool {
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		}
		return false
	}
	return helper416(nums[:len(nums)-1], target) ||
		helper416(nums[:len(nums)-1], target-nums[len(nums)-1])
}

// @lc code=end

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
