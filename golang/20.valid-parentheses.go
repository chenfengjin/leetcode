/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package main

import "fmt"

// @lc code=start
func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '[' || char == '{' || char == '(' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if !match(stack[len(stack)-1], char) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
func match(a, b byte) bool {
	if a == '{' && b == '}' {
		return true
	}
	if a == '[' && b == ']' {
		return true
	}
	if a == '(' && b == ')' {
		return true
	}
	return false
}

// @lc code=end
func main() {
	fmt.Println(isValid("()"))
}
