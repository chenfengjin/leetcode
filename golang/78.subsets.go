/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] Subsets
 */
package main

// @lc code=start
func subsets(nums []int) [][]int {
	n := 1 << len(nums)
	ret := [][]int{}
	for i := 0; i < n; i++ {
		j := i
		item := []int{}
		digit := 0
		for j > 0 {
			if j&1 == 1 {
				item = append(item, nums[digit])
			}
			digit += 1
			j >>= 1
		}
		ret = append(ret, item)
	}
	return ret
}

// @lc code=end
