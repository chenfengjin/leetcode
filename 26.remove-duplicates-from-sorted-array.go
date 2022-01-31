/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
package main

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	left := 1
	right := 1
	for ; right < len(nums); right++ {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// @lc code=end
