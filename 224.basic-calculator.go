/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
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
	sign := 1
	// remove all blank
	s = strings.Join(strings.Split(s, " "), "")
	//  small trick to deal with last operator/operand
	s = s + "#"
	for i := 1; i < len(s); i++ {
		// TODO
		//  change from digit to non digit or vice versa
		if unicode.IsDigit(rune(s[i])) != unicode.IsDigit(rune(s[i-1])) ||
			// TODO deal with this grafefully
			!unicode.IsDigit((rune(s[i]))) && !unicode.IsDigit(rune(s[i])) {
			sub := s[left:i]
			if unicode.IsDigit(rune(s[left])) { // finish reading open number
				// finish of an number
				// for example 	 1234 + 1234 + 1234 + 1234
				//						^    ^
				//						|    |
				//    			       left  i
				operand, err := strconv.Atoi(sub)
				if err != nil {
					panic(err)
				}
				{
					operands = append(operands, operand*sign)
					sign = 1
				}

				if len(operators) != 0 {
					op := operators[len(operators)-1]

					switch op {
					case "+":
						operand1 := operands[len(operands)-1]
						operand2 := operands[len(operands)-2]
						operands = operands[0 : len(operands)-2]
						operands = append(operands, operand1+operand2)
						operators = operators[0 : len(operators)-1]
					case "-":
						operand1 := operands[len(operands)-1]
						operand2 := operands[len(operands)-2]
						operands = operands[0 : len(operands)-2]
						operands = append(operands, operand2-operand1)
						operators = operators[0 : len(operators)-1]
					case "(":
					case ")":
					}
				}

			} else { // finish reading one operator
				switch string(s[left]) {
				case "+":
					fallthrough
				case "-":
					//  situation of sign character
					//  the first -1+1+3
					//  the first after '('
					//  three things about sign
					//  recognize sign character
					//  multiply sign when finish reading a number
					//  reset sign after finish reading a number
					//  a maybe solution is use strconv to deal with sign character

					// I found a graceful solution to deal with this problem
					// add zero if +/- means positive or negative
					if left == 0 || s[left-1] == byte('(') {
						// if string(s[left]) == "+" {
						// 	sign = 1
						// } else {
						// 	sign = -1
						// }
						operands = append(operands, 0)
					}
					operators = append(operators, string(s[left]))

				case "(":
					// 1234 + 1234 +1234 +	(   1234 + 1234)
					//						^	^
					// 						｜	｜
					//
					operators = append(operators, string(s[left]))
				case ")":
					operators = operators[0 : len(operators)-1]
					if len(operators) != 0 {
						op := operators[len(operators)-1]

						switch op {
						case "+":
							operand1 := operands[len(operands)-1]
							operand2 := operands[len(operands)-2]
							operands = operands[0 : len(operands)-2]
							operands = append(operands, operand1+operand2)
							operators = operators[0 : len(operators)-1]
						case "-":
							operand1 := operands[len(operands)-1]
							operand2 := operands[len(operands)-2]
							operands = operands[0 : len(operands)-2]
							operands = append(operands, operand2-operand1)
							operators = operators[0 : len(operators)-1]
						}
					}
				}

			}
			left = i
			// fmt.Println(operators)
			// fmt.Println(operands)
			// fmt.Println()
		}
	}
	return operands[0]
}

// @lc code=end
func main() {
	// fmt.Println(calculate("1+2+3+4-5"))
	// fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("-2+ 1"))

}
