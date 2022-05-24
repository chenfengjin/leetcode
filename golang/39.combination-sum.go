/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] Combination Sum
 */
package main

import "fmt"

// @lc code=start
func combinationSum(nums []int, N int) [][]int {
	used := make([]bool, len(nums))
	ret := [][]int{}
	cur := []int{}
	collect := func(item []int) {
		tmp := make([]int, len(item))
		copy(tmp, item)
		ret = append(ret, tmp)
	}
	arraySumHelper(nums, cur, used, N, 0, collect)
	return ret
}
func arraySumHelper(nums []int, cur []int, used []bool, n, K int, collect func([]int)) {
	if n < 0 {
		return
	}

	if n == 0 {
		collect(cur)
	}

	for idx := K; idx < len(nums); idx++ {
		num := nums[idx]
		cur = append(cur, num)
		used[idx] = true
		arraySumHelper(nums, cur, used, n-num, idx, collect)
		cur = cur[:len(cur)-1]
		used[idx] = false
	}
}

// @lc code=end

func main() {
	fmt.Println(ArraySum([]int{1, 3, 2, 4}, 5))
	fmt.Println(ArraySum([]int{1, 2, 3}, 3))
	fmt.Println(ArraySum([]int{1, 2, 3}, 6))
}
