/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] Combinations
 */
package main

// @lc code=start
func combine(n int, k int) [][]int {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	cur := []int{}
	ret := [][]int{}
	collect := func(item []int) {
		tmp := make([]int, len(item))
		copy(tmp, item)
		ret = append(ret, tmp)

	}
	combineHelper(nums, cur, 0, k, collect)
	return ret
}

// 排列限制了顺序就是组合
// func combineHelper(nums, cur []int, curIndex, k int, collect func([]int)) {
// 	if len(cur) == k {
// 		collect(cur)
// 		return
// 	}
// 	for i := curIndex; i < len(nums); i++ {
// 		cur = append(cur, nums[i])
// 		combineHelper(nums, cur, curIndex+1, k, collect)
// 		cur = cur[:len(cur)-1]
// 	}
// }
func combineHelper(nums, cur []int, curIndex, k int, collect func([]int)) {
	if len(cur) == k {
		collect(cur)
		return
	}
	for i := curIndex; i < len(nums); i++ {
		cur = append(cur, nums[i])
		combineHelper(nums, cur, i+1, k, collect)
		cur = cur[:len(cur)-1]
	}
}

// @lc code=end
