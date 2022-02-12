/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
package main

import (
	"fmt"
	"sort"
)

// keys
// 1. sort and then use two pointer
// 2. deal with duplication
// 3. choose a good loop condition

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))

	ret := [][]int{}
	for i := 1; i < len(nums)-1; i++ {

		left := 0
		right := len(nums) - 1

		for left < i && right > i {
			sum := nums[left] + nums[right] + nums[i]
			if sum > 0 {
				right -= 1
				continue
			} else if sum < 0 {
				left += 1
				continue
			}

			duplicated := false
			if left > 0 && nums[left] == nums[left-1] {
				duplicated = true
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				duplicated = true
			}
			if nums[i] == nums[i-1] && left != i-1 {
				duplicated = true
			}
			if !duplicated {
				ret = append(ret, []int{nums[left], nums[i], nums[right]})
			}
			//  it is maybe a good idea to simplely solve situation if sum == target
			// without using continue and break
			if right-1 > i {
				right -= 1
			} else {
				left += 1
			}

		}
	}
	return ret
}

// @lc code=end

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
