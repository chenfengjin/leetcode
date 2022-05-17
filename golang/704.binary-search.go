/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := (left + right) / 2
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	if left >= 0 && left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

// @lc code=end

