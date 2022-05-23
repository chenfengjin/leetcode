/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] Third Maximum Number
 */
package main

import (
	"container/heap"
)

// @lc code=start

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// 重新定义 less
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[1] > nums[0] {
			return nums[1]
		} else {
			return nums[0]
		}
	}

	h := IntHeap(nums)
	heap.Init(&h)

	largest := heap.Pop(&h).(int)
	second := int(0)
	third := second

	second = heap.Pop(&h).(int)
	for second == largest {
		if h.Len() == 0 {
			return largest
		}
		second = heap.Pop(&h).(int)
	}

	if h.Len() == 0 {
		return largest
	}
	third = heap.Pop(&h).(int)
	for third == second {
		if h.Len() == 0 {
			return largest
		}
		third = heap.Pop(&h).(int)

	}
	return third
}

// @lc code=end

func main() {
	thirdMax([]int{2, 2, 3, 1})
}
