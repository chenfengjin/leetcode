/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
package main

import (
	"fmt"
	"math"
)

// pay attention to conner case
// @lc code=start
func jump(nums []int) int {
	// 2,3,1,1,4
	// 0 1 1 2 2
	steps := make([]int, len(nums))
	for i := 0; i < len(steps); i++ {
		steps[i] = math.MaxInt32
	}
	if len(nums) == 1 {
		return 0
	}
	steps[0] = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		for j := 1; j <= nums[i]; j++ {
			if j+i >= len(nums) {
				break
			}
			if steps[i+j] > steps[i]+1 {
				steps[i+j] = steps[i] + 1
			}
		}
	}
	return steps[len(steps)-1]
}

// @lc code=end

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
