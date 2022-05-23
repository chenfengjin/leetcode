/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] Hamming Distance
 */
package main

// @lc code=start
func hammingDistance(x int, y int) int {
	z := x ^ y
	ret := 0
	for z > 0 {
		z &= (z - 1)
		ret++
	}
	return ret
}

// @lc code=end
