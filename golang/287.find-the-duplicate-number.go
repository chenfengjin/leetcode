/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */
package main

import "fmt"

// @lc code=start
func findDuplicate(nums []int) int {

	length := len(nums)
	if length >= 2 && nums[0] == nums[1] {
		return nums[0]
	}
	ret := 0
	for _, num := range nums {
		fmt.Println(num)
		ret ^= num
	}
	for i := 1; i < length; i++ {
		fmt.Println(i)
		ret ^= i
	}
	return ret
}

// @lc code=end

func main() {
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
