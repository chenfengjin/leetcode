/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] Single Number
 */
package main

// @lc code=start
func singleNumber(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}
	return ret
}

// @lc code=end
