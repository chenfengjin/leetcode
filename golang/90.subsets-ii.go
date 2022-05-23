/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] Subsets II
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	length := len(nums)
	ret := [][]int{}
	collect := func(item []int) {
		tmp := make([]int, len(item))
		copy(tmp, item)
		ret = append(ret, tmp)
	}
	used := make([]bool, length)
	cur := []int{}
	subsetsWithDupHelper(nums, used, cur, 0, collect)
	return ret
}

func subsetsWithDupHelper(nums []int, used []bool, cur []int, idx int, collect func([]int)) {
	collect(cur)
	for i := idx; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			if !used[i-1] {
				continue
			}
		}
		cur = append(cur, nums[i])
		used[i] = true
		subsetsWithDupHelper(nums, used, cur, i+1, collect)
		cur = cur[:len(cur)-1]
		used[i] = false
	}
}

// @lc code=end
func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 3, 3}))
}
