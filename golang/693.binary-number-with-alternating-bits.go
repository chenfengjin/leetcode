/*
 * @lc app=leetcode id=693 lang=golang
 *
 * [693] Binary Number with Alternating Bits
 */
package main

import "fmt"

// @lc code=start
func hasAlternatingBits(n int) bool {
	mask1 := 0x55555555
	mask2 := 0xAAAAAAAA
	if n == 2 {
		return true
	}
	if n&(n-1) == 0 {
		return false
	}
	return (n|mask1 == mask1 && n&mask1 == n) ||
		(n|mask2 == mask2 && n&mask2 == n)
}

// @lc code=end
func main() {
	// fmt.Println(hasAlternatingBits(5))
	// fmt.Println(hasAlternatingBits(6))
	// fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(8))
	// fmt.Println(hasAlternatingBits(9))
	// fmt.Println(hasAlternatingBits(10))
}
