/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
package main

import "fmt"

// @lc code=start
func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	cur := 0
	for cur <= right {
		if nums[cur] == 1 {
			cur++
			continue
		} else if nums[cur] == 0 {
			nums[left], nums[cur] = nums[cur], nums[left]
			left++
			cur++
		} else {
			nums[right], nums[cur] = nums[cur], nums[right]
			right--
		}
	}
}

// @lc code=end

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}

// 2,0,2,1,1,0

// 0,0,2,1,1,2

// 0,0,1,1,2,2
