/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
package main

import "fmt"

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

func trap(height []int) int {
	leftStack := NewStack()
	rightStack := NewStack()
	left := make([]int, len(height))
	right := make([]int, len(height))
	// differ from 84.largest rectangle in histogram
	// we need height
	for idx := 0; idx < len(height); idx++ {
		if leftStack.Empty() {
			left[idx] = 0
			leftStack.Push(height[idx])
			continue
		}
		if height[idx] <= height[idx-1] {
			left[idx] = leftStack.Top()
		} else {
			for !leftStack.Empty() && height[idx] > leftStack.Top() {
				leftStack.Pop()
			}
			if leftStack.Empty() {
				left[idx] = 0
				leftStack.Push(height[idx])
			} else {
				left[idx] = leftStack.Top()
			}
		}
	}
	for idx := len(height) - 1; idx >= 0; idx-- {
		if rightStack.Empty() {
			right[idx] = 0
			rightStack.Push(height[idx])
			continue
		}
		if height[idx] <= height[idx+1] {
			right[idx] = rightStack.Top()
		} else {
			for !rightStack.Empty() && height[idx] > rightStack.Top() {
				rightStack.Pop()
			}
			if rightStack.Empty() {
				right[idx] = 0
				rightStack.Push(height[idx])
			} else {
				right[idx] = rightStack.Top()
			}

		}
	}
	ret := 0
	for idx := 0; idx < len(height); idx++ {
		ret += max(min(left[idx], right[idx])-height[idx], 0)

	}
	return ret
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
