/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */
package main

import "fmt"

// @lc code=start
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 {
			// coner case 2
			continue
		}
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i+1 {
			// coner case 1
			if nums[i] == nums[nums[i]-1] {
				break
			}
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]

			nums[i] = tmp
		}
	}

	// fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

// @lc code=end
func main() {
	nums := []int{1, 2, 0}
	nums = []int{3, 4, -1, 1}
	// nums = []int{7, 8, 9, 11, 12}
	// nums = []int{1, 2, 6, 3, 5, 4}
	nums = []int{2, 2}
	nums = []int{1, 2, 6, 3, 5, 4}
	fmt.Println(firstMissingPositive(nums))
}
