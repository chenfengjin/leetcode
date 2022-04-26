/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */
package main

import (
	"fmt"
	"strings"
)

// @lc code=start
func convertToTitle(columnNumber int) string {
	ret := []string{}
	columnNumber = columnNumber
	for columnNumber > 0 {
		ret = append(ret, string([]byte{byte(columnNumber%26 + 'A')}))
		columnNumber /= 26
	}
	return strings.Join(ret, "")
}

// @lc code=end
func main() {
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(701))
}
