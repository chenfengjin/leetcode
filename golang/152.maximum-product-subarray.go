/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */
package main

import "fmt"

// @lc code=start
func maxProduct(nums []int) int {
	length := len(nums)
	dpMax := make([]int, length)
	dpMin := make([]int, length)
	copy(dpMax, nums)
	copy(dpMin, nums)
	for i := 1; i < length; i++ {
		num := nums[i]
		if dpMin[i] > dpMin[i-1]*nums[i] {
			dpMin[i] = dpMin[i-1] * nums[i]
		}
		if dpMin[i] > dpMax[i-1]*num {
			dpMin[i] = dpMax[i-1] * num
		}

		if dpMax[i] < dpMax[i-1]*num {
			dpMax[i] = dpMax[i-1] * num
		}
		if dpMax[i] < dpMin[i-1]*nums[i] {
			dpMax[i] = dpMin[i-1] * nums[i]
		}
	}
	max := dpMax[0]
	for i := 0; i < len(dpMax); i++ {
		if max < dpMax[i] {
			max = dpMax[i]
		}
	}
	return max
	// return dpMax[len(dpMax)-1]
}

// @lc code=end

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	// fmt.Println(maxProduct([]int{-2, 0, -1}))
}
