/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */
package main

import (
	"fmt"
	"sort"
)

// key is how to deal with duplicates
//  we can add duplicates in order, if element[i] is equal to element[i-1] && element[i] is not selected
//  element[i] shoud neither be selected
// @lc code=start
func subsetsWithDup(nums []int) [][]int {

	sort.IntSlice(nums).Sort()

	n := 1 << len(nums)
	ret := [][]int{[]int{}}
	for i := 1; i < n; i++ {
		item := []int{}
		m := 0
		j := i
		found := true
		for j > 0 {
			if j&1 == 1 {
				if m >= 1 && nums[m] == nums[m-1] && (i>>(m-1))&1 != 1 {
					found = false
				}
				item = append(item, nums[m])
			}
			j >>= 1
			m++
		}
		if len(item) != 0 && found {
			ret = append(ret, item)
		}
	}
	return ret
}

// @lc code=end
func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

// 0 1 1
// 1 0 1
