/*
 * @lc app=leetcode id=393 lang=golang
 *
 * [393] UTF-8 Validation
 */
// pay attention to conner case
// 1. larer than 8
package main

import "fmt"

// @lc code=start
func validUtf8(data []int) bool {
	if len(data) == 0 || len(data) > 4 {
		return false
	}
	length := len(data)
	if length == 1 {
		return data[0]>>7|1 == 0
	}

	n := 0
	mask := 0
	for ; n < length; n++ {
		//  order of this two lines is important
		mask <<= 1
		mask |= 1
	}

	for ; n < 8; n++ {
		//  shound not shift left mask
		data[0] >>= 1
	}
	if data[0]&mask != data[0] {
		return false
	}
	// for i := 1; i < length; i++ {
	if data[1]>>6 != 2 {
		return false
	}
	// }
	return true
}

// @lc code=end

func main() {
	// fmt.Println(validUtf8([]int{197, 130, 3, 1}))
	fmt.Printf("%b", 140)
}
