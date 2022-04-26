/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */
package main

// @lc code=start
func canJump(nums []int) bool {
	// 2,3,1,1,4
	// bool 
	// x x x x x 
	length := len(nums)
	max := 0
	reachable := make([]int, length)

	for i := 0; i < len(nums); i++ {
		if i+nums[i] >= length {
			return true
		}
	}

}

// @lc code=end
