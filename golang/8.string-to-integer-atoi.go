/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
package main

import "fmt"

// @lc code=start
func myAtoi(s string) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		ret *= 10
		ret += int(s[i] - '0')
	}

}

// @lc code=end

func main() {
	fmt.Println(myAtoi("1111111111111"))
}
