/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */
package main

import "sort"

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.IntSlice(nums).Sort()

	used := make([]bool, len(nums))
	current := []int{}
	result := [][]int{}
	helper47(nums, current, used, &result)
	return result
}

// 副本中的array指针与原slice指向同一个地址，所以当修改副本slice的元素时，原slice的元素值也会被修改。但是如果修改的是副本slice的len和cap时，原slice的len和cap仍保持不变
// 如果在操作副本时由于扩容操作导致重新分配了副本slice的array内存地址，那么之后对副本slice的操作则完全无法影响到原slice，包括slice中的元素

func helper47(nums []int, current []int, used []bool, results *[][]int) {
	if len(current) == len(nums) {
		tmp := []int{}
		for _, n := range current {
			tmp = append(tmp, n)
		}
		*results = append(*results, tmp)
		return
	}

	for idx, num := range nums {
		if idx > 0 && nums[idx] == nums[idx-1] && !used[idx-1] {
			continue
		}
		if !used[idx] {
			current = append(current, num)
			used[idx] = true
			helper47(nums, current, used, results)
			current = current[0 : len(current)-1]
			used[idx] = false
		}
	}
}

// @lc code=end
