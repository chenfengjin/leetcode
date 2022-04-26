/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
package main

// @lc code=start
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) // 1
	for left < right { // 2
		middle := (left + right) / 2 //
		if nums[middle] < target {   // 4
			left = middle + 1 //5
		} else {
			right = middle
		}
	}
	return left // 6
}

// @lc code=end
