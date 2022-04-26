/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */
package main

import "fmt"

// TODO  可以通过单调堆来优化常数空间
// @lc code=start
func lengthOfLIS(nums []int) int {
	// return lengthOfLISDP(nums)
	return lengthOfLISWithStack(nums)
}

func lengthOfLISGreedy(nums []int) int {
	return 0
}

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

func lengthOfLISWithStack(nums []int) int {
	stack := NewStack()
	if len(nums) == 0 {
		return 0
	}

	dp := []int{1}
	stack.Push(0)
	max := 1

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		current := 1

		if nums[stack.Top()] < num {
			stack.Push(i)
			current = dp[i-1] + 1
		} else {
			for !stack.Empty() && nums[stack.Top()] >= num {
				stack.Pop()
			}
			if stack.Empty() {
				stack.Push(i)
				current = 1
			} else {
				current = dp[stack.Top()] + 1
				stack.Push(i)
			}
		}
		dp = append(dp, current)
		if max < current {
			max = current
		}
	}
	fmt.Println(dp)

	return max
}
func lengthOfLISDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := []int{1}
	max := 1
	for i, num := range nums[1:] {
		current := 1
		for j := 0; j < i+1; j++ {
			if nums[j] < num && dp[j]+1 > current {
				current = dp[j] + 1
			}
		}
		dp = append(dp, current)
		if current > max {
			max = current
		}
	}

	return max
}

// @lc code=end
func main() {
	// fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))

}

//   10,9,2,5,3,7,101,18
//  stack 10  	1   1
//  stack 9  	1   1,1
//  stack 2		1   1,1,1
//  stack 2 5   2 	1,1,1,2
//  stack 2 3   2	1,1,1,2,2
