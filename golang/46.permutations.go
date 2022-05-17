/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] Permutations
 */
package main

// @lc code=start
func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	cur := []int{}
	ret := [][]int{}
	collector := func(item []int) {
		// 这里会被编译器优化
		// tmpL=item[:]

		tmp := make([]int, len(item))
		copy(tmp, item)
		ret = append(ret, tmp)
	}
	permuteHelper(nums, used, cur, collector)
	return ret
}

func permuteHelper(nums []int, used []bool, cur []int, collector func([]int)) {
	if len(cur) == len(nums) {
		collector(cur)
		return
	}
	for idx, num := range nums {
		if !used[idx] {
			cur = append(cur, num)
			used[idx] = true
			permuteHelper(nums, used, cur, collector)
			cur = cur[:len(cur)-1]
			used[idx] = false
		}
	}
}

// @lc code=end
