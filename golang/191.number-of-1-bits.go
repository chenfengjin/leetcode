/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */
package main

// @lc code=start
func hammingWeight(num uint32) int {
	ret := 0
	for num > 0 {
		num = num & (num - 1)
		ret++
	}
	return ret
}

// @lc code=end

// 1010101 -> 4
//  101111 -> 5
