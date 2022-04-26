/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */
package main

import "fmt"

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prefixSum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] == 1 {
			prefixSum = append(prefixSum, prefixSum[i-1]+1)
		} else {
			prefixSum = append(prefixSum, 0)
		}
	}
	max := 0
	for i := 0; i < len(prefixSum); i++ {
		if max < prefixSum[i] {
			max = prefixSum[i]
		}
	}
	return max
}

// @lc code=end

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1}))
}
