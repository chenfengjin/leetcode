/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
package main

// @lc code=start
func mySqrt(x int) int {
	var right int
	if x < 10000 {
		right = x / 2
	} else if x < 1000000 {
		right = x / 100
	} else {
		right = x / 1000
	}
	_ = right
	return 0
	// left := 0
	// right := x / 100
	// x 10000 100000000

	// sqrt(x) < x/2
	// sqrt(x) < x/100
}

// @lc code=end
