/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */
package main

import (
	"fmt"
	"strconv"
)

// key is how to deal with zero
//  TODO code could be dealed with carefully
// @lc code=start
func restoreIpAddresses(s string) []string {
	return helper93(s, 4)
}
func helper93(s string, count int) []string {
	if len(s) < count {
		return []string{}
	}
	if count == 1 {
		if len(s) > 1 && s[0] == '0' {
			return []string{}
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return []string{}
		}
		if n < 0 {
			return []string{}
		}
		if n > 255 {
			return []string{}
		}
		return []string{s}
	}
	ret := []string{}
	for i := 1; i < 4; i++ {
		if i > len(s) {
			break
		}
		if i > 1 && s[0] == '0' {
			break
		}
		cur, _ := strconv.Atoi(s[0:i])
		if cur >= 0 && cur <= 255 {
			for _, sub := range helper93(s[i:], count-1) {
				ret = append(ret, s[0:i]+"."+sub)
			}
		}

	}
	return ret
}

// @lc code=end
func main() {
	fmt.Println(strconv.Atoi(""))

}
