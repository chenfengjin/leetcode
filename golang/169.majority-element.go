/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */
package main

import "fmt"

// @lc code=start
func majorityElement(nums []int) int {
	ret := nums[0]
	count := 1
	for _, num := range nums[1:] {
		if num == ret {
			count += 1
		} else {
			if count > 0 {
				count -= 1
			} else {
				count = 1
				ret = num

			}
		}
	}
	return ret
}

// @lc code=end
func main() {
	fmt.Println(majorityElement([]int{1, 3, 2, 2}))
	fmt.Println(majorityElement([]int{1, 3, 4, 2, 2}))
}
