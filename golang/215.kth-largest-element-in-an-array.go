/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */
package main

import "container/heap"

// @lc code=start
type IntHeap []int

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func findKthLargest(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)
	ret := 0
	for i := 0; i < k; i++ {
		ret = heap.Pop(&h).(int)
	}
	return ret
}

// @lc code=end
