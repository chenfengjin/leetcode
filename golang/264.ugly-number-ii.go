/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */
package main

// @lc code=start

type IntHeap []int

var (
// _ heap.Interface = &IntHeap{}
)

// func (h IntHeap) Len() int {
// 	return len(h)
// }
// func (h IntHeap) Less(i, j int) bool {
// 	return h[i] < h[j]
// }
// func (h IntHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[j]
// }
// func (h *IntHeap) Push(val interface{}) {
// 	*h = append(*h, val.(int))
// }
// func (h *IntHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// func nthUglyNumber(n int) int {
// 	seen := map[int]bool{}
// 	h := &IntHeap{1}
// 	heap.Init(h)

// 	for i := 1; ; i++ {
// 		if i == n {
// 			return heap.Pop(h).(int)
// 		}
// 		x := h.Pop().(int)

// 	}
// }

// @lc code=end
