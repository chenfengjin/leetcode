/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start

type IntervalSlice [][]int

func (intervals IntervalSlice) Len() int {
	return len(intervals)

}
func (interval IntervalSlice) Less(i, j int) bool {
	return interval[i][0] < interval[j][0]
}

func (intervals *IntervalSlice) Swap(i, j int) {
	(*intervals)[i], (*intervals)[j] = (*intervals)[j], (*intervals)[i]
}
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	intervalSlice := IntervalSlice(intervals)
	sort.Sort(&intervalSlice)

	ret := [][]int{}
	cur := intervalSlice[0]
	for _, interval := range intervalSlice[1:] {
		if interval[0] <= cur[1] {
			if interval[1] > cur[1] {
				cur[1] = interval[1]
			}
		} else {
			ret = append(ret, cur)
			cur = interval
		}
	}
	ret = append(ret, cur)
	return ret

}

// @lc code=end

func main() {
	fmt.Println(merge(([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})))
}
