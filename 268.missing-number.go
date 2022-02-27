/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */
package main

import "fmt"

//  bit-manipulation
// @lc code=start
func missingNumber(nums []int) int {
	n := len(nums)
	ret := 0
	for i := 1; i <= n; i++ {
		ret ^= i
	}
	for _, num := range nums {
		ret ^= num
	}
	return ret

}

// @lc code=end
func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}
