/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package main

// @lc code=start
func isValid(s string) bool {
	stack := []string{}
	parentheses := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	i := 0

	for ; i < len(s); i++ {
		char := s[i]
		switch string(char) {
		case "[":
			fallthrough
		case "(":
			fallthrough
		case "{":
			stack = append(stack, string(char))
		case "}":
			fallthrough
		case "]":
			fallthrough
		case ")":
			if len(stack) == 0 {
				return false
			}
			if parentheses[string(char)] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0

}

// @lc code=end


7 2 4 3 
  5 6 4 