/*
 * @lc app=leetcode id=136 lang=golang
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

// 交换律
// 结合律
//  a^a=0
// 1^2^3^4^3^1^4
// 1^1^3^3^4^4^2
// @lc code=end
