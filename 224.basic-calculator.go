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

// some key point
// 1. how to deal with ()
// 2. how to deal with -

// a small trick
// just regard + and - as sign rather than plus/minus
// add all numbers works
// @lc code=start
func calculate(s string) int {
	return calculateUsingStack(s)
}
func calculate1(s string) int {
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
					//  is op is + or -, look ahead
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
						{
							// loop is not neccessary as there is only one element in op stack
							if len(operators) > 0 {
								operator := operators[len(operators)-1]
								if string(operator) == "+" || string(operator) == "-" {
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
						{

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
	return operands[0]
}

type stack struct {
	elem      []int
	operators *OpeartorStack
}

func (s *stack) Push(elem int) {
	s.elem = append(s.elem, elem)
}

// func (s *stack) pushOperator(string) {
//
// }
func (s *stack) Size() int {
	return len(s.elem)
}
func (s *stack) Execute() {
	operand1 := s.elem[len(s.elem)-2]
	operand2 := s.elem[len(s.elem)-1]
	s.elem = s.elem[0 : len(s.elem)-2]
	top := 0
	op := s.operators.Pop()
	switch op {
	case "+":
		top = operand1 + operand2
	case "-":
		top = operand1 - operand2
	case "/":
		top = operand1 / operand2
	case "*":
		top = operand1 / operand2
	}
	s.Push(top)
}
func (s *stack) Top() int {
	return s.elem[len(s.elem)-1]
}

type OpeartorStack struct {
	elemtns []string
}

func (s *OpeartorStack) Push(op string) {
	s.elemtns = append(s.elemtns, op)
}
func (s *OpeartorStack) Top() string {
	return s.elemtns[len(s.elemtns)-1]
}
func (s *OpeartorStack) Pop() string {
	op := s.elemtns[len(s.elemtns)-1]
	s.elemtns = s.elemtns[0 : len(s.elemtns)-1]
	return op

}
func (s *OpeartorStack) Empty() bool {
	return len(s.elemtns) == 0
}

func newStack(operators *OpeartorStack) *stack {
	return &stack{
		operators: operators,
	}
}
func calculateUsingStack(s string) int {
	if len(s) == 1 {
		return int(s[0] - byte('0'))
	}
	ops := &OpeartorStack{}
	stack := newStack(ops)
	left := 0
	right := 1
	// remove all space
	s = strings.Join(strings.Split(s, " "), "")
	s = s + "#"
	for ; right < len(s); right++ {
		if unicode.IsDigit(rune(s[right])) && unicode.IsDigit(rune(s[right-1])) {
			continue
		}

		if unicode.IsDigit(rune(s[left])) { // 找到了数字
			operand, _ := strconv.Atoi(s[left:right])
			stack.Push(operand)
			if stack.Size() >= 2 && ops.Top() != "(" {
				stack.Execute()
			}
		} else { // 找到了符号
			op := s[left:right]
			switch op {
			case "+":
				ops.Push(op)
			case "-":
				if left == 0 || s[left-1:left] == "(" {
					stack.Push(0)
				}
				ops.Push(op)

			case "(":
				ops.Push(op)
			case ")":
				ops.Pop()
				if stack.Size() >= 2 {
					stack.Execute()
				}
			case "#":
				if stack.Size() >= 2 {
					stack.Execute()
				}
			default:
			}
		}
		left = right
	}
	return stack.Top()
}

// @lc code=end
func main() {
	// fmt.Println(calculateUsingStack("1+2+3+4"))
	// fmt.Println(calculateUsingStack("1+2-3+4"))
	fmt.Println(calculateUsingStack("1+(3-(2-1))+4"))

	// fmt.Println(calculate("1+2+3+4-5"))
	// fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	// fmt.Println(calculate("-2+ 1"))

	// fmt.Println(calculate("3+4-5*2"))

}
