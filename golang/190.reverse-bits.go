/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */
package main

// @lc code=start
func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			ret <<= 1
			// ret += 1
			ret |= 1
		} else {
			ret <<= 1
		}
		num >>= 1
	}
	return ret
}

// @lc code=end
// 0111
// 111
// 0001

// 1000
