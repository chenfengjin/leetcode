/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] Combination Sum III
 */
package main

import "fmt"

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	cur := []int{}
	m := map[int][][]int{}
	collect := func(item []int, key int) {
		tmp := make([]int, k)
		copy(tmp, item)
		m[key] = append(m[key], tmp)
	}
	combinationSum3Helper(nums, cur, k, 0, 0, collect)
	for key, item := range m {
		fmt.Println(key, item)
	}
	return m[n]
}
func combinationSum3Helper(nums, cur []int, k, curSum int, idx int, collect func([]int, int)) {
	if len(cur) == k {
		collect(cur, curSum)
	}
	for i := idx; i < len(nums); i++ {
		num := nums[i]
		cur = append(cur, num)
		combinationSum3Helper(nums, cur, k, curSum+num, i+1, collect)
		cur = cur[:len(cur)-1]
	}
}

// @lc code=end
func main() {
	fmt.Println(combinationSum3(4, 1))
}
