/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] Simplify Path
 */
package main

import "strings"

// @lc code=start
type stack struct {
	data []string
}

func newStack() *stack {
	return &stack{}
}
func (s *stack) Push(item string) {
	s.data = append(s.data, item)
}
func (s *stack) Pop() {
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func simplifyPath(path string) string {
	s := newStack()
	pathArray := strings.Split(path, "/")
	for _, item := range pathArray {
		// for //
		if item == "" || item == "." {
			continue
		}
		if item == ".." {
			s.Pop()
		} else {
			s.Push(item)
		}

	}
	return "/" + strings.Join(s.data, "/")
}

// @lc code=end
