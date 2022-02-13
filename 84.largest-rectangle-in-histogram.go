/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
package main

import "fmt"

// Monotonic stack is a good choice for this kind of problem
// an import key is choosing direction of monotonic stack

// there are three keys
// 1. direction of monotonic stack, it is determined by what areas array mean
//  increasing  <==> max area of rectangle that ends with current index
//  decreasing  <==> max area of rectangle that starts with this the first index in stack
// 2. what to do when pop from stack
// 3. what to store in stack, stack element should be index rather than heights

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

//  you can only push or pop
//  but you can read all the elements of an stack data structure
func (s *Stack) Elems() []int {
	return s.elems
}
func largestRectangleArea(heights []int) int {
	leftStack := NewStack()
	rightStack := NewStack()
	left := make([]int, len(heights))
	right := make([]int, len(heights))
	for idx := 0; idx < len(heights); idx++ {
		height := heights[idx]
		if leftStack.Empty() {
			left[idx] = -1
			leftStack.Push(idx)
			continue
		}
		if heights[leftStack.Top()] < height {
			// TODO nothing to do, not neccessary to add something
			// TODO when equal
			left[idx] = idx - 1
			leftStack.Push(idx)
		} else {
			for !leftStack.Empty() && heights[leftStack.Top()] >= height {
				leftStack.Pop()
			}
			if leftStack.Empty() {
				left[idx] = -1
			} else {
				left[idx] = leftStack.Top()
			}
			leftStack.Push(idx)

		}
	}
	// fmt.Println(left)
	for idx := len(heights) - 1; idx >= 0; idx-- {
		if rightStack.Empty() {
			right[idx] = len(heights)
			rightStack.Push(idx)
			continue
		}
		height := heights[idx]
		if height > heights[idx+1] {
			right[idx] = idx + 1
			rightStack.Push(idx)
		} else {
			for !rightStack.Empty() && height <= heights[rightStack.Top()] {
				rightStack.Pop()
			}
			if rightStack.Empty() {
				right[idx] = len(heights)
			} else {
				right[idx] = rightStack.Top()
			}
			rightStack.Push(idx)
		}
	}
	// fmt.Println(right)

	max := 0
	for i := 0; i < len(heights); i++ {
		cur := heights[i] * (right[i] - left[i] - 1)
		if max < cur {
			max = cur
		}
	}
	return max
}

// @lc code=end

func main() {
	fmt.Println(largestRectangleArea([]int{1, 1}))
}
