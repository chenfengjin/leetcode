/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

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
	// remove all space
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
					// is op is * or /,do nothing
					// is op is + or -, look ahead
					case "+":
						//  add # suffix is realy a good choice
						if string(s[i]) == "*" || string(s[i]) == "/" {
							//Dothing if next operator has high priority
						} else {
							operand1 := operands[len(operands)-1]
							operand2 := operands[len(operands)-2]
							operands = operands[0 : len(operands)-2]
							operands = append(operands, operand1+operand2)
							operators = operators[0 : len(operators)-1]
						}
					case "-":
						if string(s[i]) == "*" || string(s[i]) == "/" {
							//Dothing if next operator has high priority
						} else {
							operand1 := operands[len(operands)-1]
							operand2 := operands[len(operands)-2]
							operands = operands[0 : len(operands)-2]
							operands = append(operands, operand2-operand1)
							operators = operators[0 : len(operators)-1]
						}
					case "*":
						operand1 := operands[len(operands)-1]
						operand2 := operands[len(operands)-2]
						operands = operands[0 : len(operands)-2]
						operands = append(operands, operand1*operand2)
						operators = operators[0 : len(operators)-1]
						if false {
							// loop is not neccessary as there is at least one element in op stack
							if len(operators) > 0 {
								operator := operators[len(operators)-1]
								if string(operator) == "+" || string(operator) == "-" &&
									string(s[i]) != "*" && string(s[i]) != "/" {
									operators = operators[0 : len(operators)-1]

									operand1 := operands[len(operands)-1]
									operand2 := operands[len(operands)-2]
									operands = operands[0 : len(operands)-2]
									if operator == "+" {
										operands = append(operands, operand2+operand1)
									} else {
										operands = append(operands, operand2-operand1)
									}
								}
							}
						}
					case "/":
						operand1 := operands[len(operands)-1]
						operand2 := operands[len(operands)-2]
						operands = operands[0 : len(operands)-2]
						operands = append(operands, operand2/operand1)
						operators = operators[0 : len(operators)-1]
						if false {
							// loop is not neccessary as there is at least one element in op stack
							if len(operators) > 0 {
								operator := operators[len(operators)-1]
								if string(operator) == "+" || string(operator) == "-" &&
									string(s[i]) != "*" && string(s[i]) != "/" {
									operators = operators[0 : len(operators)-1]

									operand1 := operands[len(operands)-1]
									operand2 := operands[len(operands)-2]
									operands = operands[0 : len(operands)-2]
									if operator == "+" {
										operands = append(operands, operand2+operand1)
									} else {
										operands = append(operands, operand2-operand1)
									}
								}
							}
						}
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
					// before push operator to operator stack,make sure there is no operator in the same priority
					// loop is not neccessary
					{
						if len(operators) > 0 {
							op := operators[len(operators)-1]
							if op == "+" {
								operand1 := operands[len(operands)-1]
								operand2 := operands[len(operands)-2]
								operands = operands[0 : len(operands)-2]
								operands = append(operands, operand1+operand2)
								operators = operators[0 : len(operators)-1]

							} else if op == "-" {
								operand1 := operands[len(operands)-1]
								operand2 := operands[len(operands)-2]
								operands = operands[0 : len(operands)-2]
								operands = append(operands, operand2-operand1)
								operators = operators[0 : len(operators)-1]
							}
						}
					}
					operators = append(operators, string(s[left]))
				case "*":
					fallthrough
				case "/":
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
	// TODO 确认下出现这种情况的条件
	if len(operators) > 0 {
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
	return operands[0]
}

// @lc code=end

func main() {
	// fmt.Println(calculate("1+2+3+4-5"))
	// fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	// fmt.Println(calculate("-2+ 1"))

	// fmt.Println(calculate("3+4-5*2"))
	// fmt.Println(calculate("3+5/2"))
	// fmt.Println(calculate("1+2*5/3+6/4*2"))
	fmt.Println(calculate("1*2-3/4+5*6-7*8+9/10"))

}
