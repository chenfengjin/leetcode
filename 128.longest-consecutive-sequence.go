/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */
package main

import "fmt"

//  TODO
// 1. 路径压缩
// 2. 添加 union 时的 count 记录
// 3. 完全可以重写一遍

// @lc code=start
func union(parents map[int]int, x, y int) {
	rootX, _ := find(parents, x)
	rootY, _ := find(parents, y)
	if rootX == rootY {
		return
	}
	parents[rootY] = rootX
}
func find(parents map[int]int, elem int) (parent int, exist bool) {
	if _, ok := parents[elem]; !ok {
		return 0, false
	}
	for parents[elem] != elem {
		elem = parents[elem]
	}
	return elem, true
}

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	parents := map[int]int{}
	for _, num := range nums {
		parents[num] = num
	}
	for _, num := range nums {
		parent, exist := find(parents, num+1)
		if exist {
			union(parents, parent, num)
		}
	}
	max := 1
	counts := map[int]int{}
	for _, a := range parents {
		parent, _ := find(parents, a)
		if _, ok := counts[parent]; !ok {
			counts[parent] = 1
		} else {
			counts[parent] += 1
		}
	}
	for _, count := range counts {
		if count > max {
			max = count
		}
	}
	return max
}

// @lc code=end

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
