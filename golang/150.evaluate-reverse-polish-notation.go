/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */
package main

import (
	"fmt"
	"strconv"
)

// @lc code=start

type Stack struct {
	elems []int
}

func NewStack() *Stack {
	return &Stack{}
}
func (s *Stack) Push(n int) {
	s.elems = append(s.elems, n)
}
func (s *Stack) Pop() int {
	//
	n := s.elems[len(s.elems)-1]
	s.elems = s.elems[0 : len(s.elems)-1]
	return n
}
func (s *Stack) Top() int {
	return s.elems[len(s.elems)-1]
}
func (s *Stack) Size() int {
	return len(s.elems)
}
func (s *Stack) Empty() bool {
	return len(s.elems) == 0
}

func evalRPN(tokens []string) int {
	operators := map[string]func(a, b int) int{
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
		"*": func(a, b int) int {
			return a * b
		},
		"/": func(a, b int) int {
			return a / b
		},
	}
	stack := NewStack()
	for _, token := range tokens {
		f, ok := operators[token]
		if ok {
			operand1 := stack.Pop()
			operand2 := stack.Pop()
			result := f(operand2, operand1)
			stack.Push(result)
		} else {
			operand, _ := strconv.Atoi(token)
			stack.Push(operand)
		}
	}
	return stack.Top()
}

// @lc code=end
func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))

}
