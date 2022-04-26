/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
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

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := NewStack()
	m := map[int]int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		if stack.Empty() {
			stack.Push(num)
			m[num] = -1
		} else {
			for !stack.Empty() && num >= stack.Top() {
				stack.Pop()
			}
			if stack.Empty() {
				stack.Push(num)
				m[num] = -1
			} else {
				m[num] = stack.Top()
				stack.Push(num)
			}
		}
	}
	ret := []int{}
	for _, num := range nums1 {
		ret = append(ret, m[num])
	}
	return ret
}

// @lc code=end
func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
