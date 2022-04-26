/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */
package main

import "sort"

// differ from and two sum, sort may be a efficient way as sort is an O(nlog(n)) operation
// the same for three sum
// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Sort(sort.IntSlice(nums))

}

// @lc code=end
