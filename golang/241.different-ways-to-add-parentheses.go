/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 */
package main

import "strconv"

// @lc code=start
func diffWaysToCompute(expression string) []int {
	if isDigital(expression) {
		num, _ := strconv.ParseInt(expression, 10, 32)
		return []int{int(num)}
	}
	idx := 0
	ret := []int{}
	functions := map[byte]func(a, b int) int{
		'+': func(a, b int) int { return a + b },
		'-': func(a, b int) int { return a - b },
		'*': func(a, b int) int { return a * b },
	}
	for idx < len(expression) {
		char := expression[idx]
		if isOperator(char) {
			leftResult := diffWaysToCompute(expression[:idx])
			rightResults := diffWaysToCompute(expression[idx+1:])
			function := functions[char]
			for _, left := range leftResult {
				for _, right := range rightResults {
					ret = append(ret, function(left, right))
				}
			}
		}
		idx++
	}
	return ret
}
func isDigital(s string) bool {
	for _, char := range s {
		if char > '9' || char < '0' {
			return false
		}
	}
	return true
}
func isOperator(char byte) bool {
	return char == '+' || char == '-' || char == '*'
}

// @lc code=end
