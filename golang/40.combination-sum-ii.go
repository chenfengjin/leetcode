/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] Combination Sum II
 */
package main

import (
	"sort"
)

// @lc code=start
func combinationSum2(nums []int, N int) [][]int {
	used := make([]bool, len(nums))
	ret := [][]int{}
	cur := []int{}
	collect := func(item []int) {
		tmp := make([]int, len(item))
		copy(tmp, item)
		ret = append(ret, tmp)
	}
	sort.Sort(sort.IntSlice(nums))
	arraySumHelper2(nums, cur, used, N, 0, collect)
	return ret

}

func arraySumHelper2(nums []int, cur []int, used []bool, n, K int, collect func([]int)) {
	if n < 0 {
		return
	}

	if n == 0 {
		collect(cur)
	}

	for idx := K; idx < len(nums); idx++ {
		num := nums[idx]
		if idx > 0 && nums[idx] == nums[idx-1] && !used[idx-1] {
			continue
		}
		if !used[idx] {
			cur = append(cur, num)
			used[idx] = true
			arraySumHelper2(nums, cur, used, n-num, idx+1, collect)
			cur = cur[:len(cur)-1]
			used[idx] = false
		}
	}
}

// @lc code=end
