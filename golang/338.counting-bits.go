/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */
package main

import "fmt"

// @lc code=start
func countBits(n int) []int {
	hightBit := 0
	ret := []int{0}
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			hightBit = i
			ret = append(ret, 1)
		} else {
			ret = append(ret, ret[i-hightBit]+1)
		}
	}
	return ret
}

// @lc code=end

func main() {
	fmt.Println(countBits(5))
}
