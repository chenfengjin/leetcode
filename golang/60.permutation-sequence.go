/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */
package main

// @lc code=start
// this is a math problem
// it can be sovled by math method
// the max number of k permutation is at most n!
//
func getPermutation(n int, k int) string {
	// zero is useless
	limits := []int{0, 1}
	for i := 2; i < 10; i++ {
		limits = append(limits, limits[len(limits)-1]*i)
	}
	if n > limits[k] {
		return ""
	}
	for n > 0 {
		f
	}
}

// @lc code=end
