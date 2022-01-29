/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */
package main

//
// @lc code=start
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			break
		}
	}
	if i == len(nums)-1 {
		return
	}
	j := i + 1
	for ; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end
