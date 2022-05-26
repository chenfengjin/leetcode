/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */
package main

import "fmt"

// @lc code=start
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	prefixSum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		prefixSum = append(prefixSum, prefixSum[i-1]*num)
	}
	suffixSum := make([]int, len(nums))
	suffixSum[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] * nums[i]
	}
	ret := make([]int, len(nums))
	ret[0] = suffixSum[1]
	ret[len(nums)-1] = prefixSum[len(nums)-1-1]

	for i := 1; i < len(nums)-1; i++ {
		ret[i] = prefixSum[i-1] * suffixSum[i+1]
	}
	return ret
}

// @lc code=end

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
