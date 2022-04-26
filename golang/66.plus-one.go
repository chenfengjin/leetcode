/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
package main

import "fmt"

// import "fmt"

// @lc code=start
// 高精度数
func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i > 0; i-- {
		digits[i] = (digits[i] + carry) % 10
		carry = (digits[i] + 1) / 10
	}
	results := []int{}
	// fmt.Println("carry", carry)
	if carry == 1 {
		results = append(results, 1)
	}
	fmt.Println(digits)
	results = append(results, digits...)
	return results
}

//     1 2 3
// +       1
// -------0---

// @lc code=end

func main() {
	fmt.Println(plusOne([]int{0}))
}
