/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 */
package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// a small trick
// just regard + and - as sign rather than plus/minus
// add all numbers works
// @lc code=start
func calculate(s string) int {
	//  先去掉所有空格
	// 	转换成 string 的数组
	if len(s) == 1 {
		return int(s[0] - byte('0'))
	}
	operands := []int{}
	operators := []string{}
	left := 0
	//  small trick to deal with last operator/operand
	s = s + "#"
	for i := 1; i < len(s); i++ {
		// TODO
		//  change from digit to non digit or vice versa
		if unicode.IsDigit(rune(s[i])) != unicode.IsDigit(rune(s[i-1])) {
			fmt.Println(operators)
			fmt.Println(operands)
			sub := s[left:i]
			if unicode.IsDigit(rune(s[left])) {
				// finish of an number
				// for example 	 1234 + 1234 + 1234 + 1234
				//						^    ^
				//						|    |
				//    			       left  i
				operand, _ := strconv.Atoi(sub)
				operands = append(operands, operand)

				if len(operators) != 0 {
					op := operators[len(operators)-1]
					operand1 := operands[len(operands)-1]
					operand2 := operands[len(operands)-2]
					operands = operands[0 : len(operands)-2]

					switch op {
					case "+":
						operands = append(operands, operand1+operand2)
					case "-":
						operands = append(operands, operand2-operand1)
					}
				}

			} else {
				switch string(s[left]) {
				case "+":
					fallthrough
				case "-":
					operators = append(operators, string(s[left]))

				}

			}
			left = i
		}
	}
	return operands[0]
}

// @lc code=end
func main() {
	fmt.Println(calculate("1+2+3+4-5"))
}
