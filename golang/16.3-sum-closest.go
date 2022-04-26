/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
package main

import "sort"

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	ret := nums[0] + nums[1] + nums[2]
	sort.IntSlice(nums).Sort()
	for i := 1; i < len(nums)-1; i++ {
		left := 0
		right := len(nums) - 1
		diff := nums[i] + nums[left] + nums[right] - target

		for {
			if diff == 0{
				return target
			}
			if diff > 0 && nums[i]+nums[left]+nums[right-1]-target > 0 {
				right = right - 1

			}else if diff <0 && nums[i]+nums[left+1]+nums[right]-target 
		}
	}
}

// @lc code=end
