/*
 * @lc app=leetcode id=1723 lang=golang
 *
 * [1723] Find Minimum Time to Finish All Jobs
 */
package main

import "sort"

// @lc code=start
func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	l, r := jobs[0], 0
	for _, v := range jobs {
		r += v
	}
	return l + sort.Search(r-l, func(limit int) bool {
		limit += l
		workloads := make([]int, k)
		var backtrack func(int) bool
		backtrack = func(idx int) bool {
			if idx == n {
				return true
			}
			cur := jobs[idx]
			for i := range workloads {
				if workloads[i]+cur <= limit {
					workloads[i] += cur
					if backtrack(idx + 1) {
						return true
					}
					workloads[i] -= cur
				}
				// 如果当前工人未被分配工作，那么下一个工人也必然未被分配工作
				// 或者当前工作恰能使该工人的工作量达到了上限
				// 这两种情况下我们无需尝试继续分配工作
				if workloads[i] == 0 || workloads[i]+cur == limit {
					break
				}
			}
			return false
		}
		return backtrack(0)
	})

	// 作者：LeetCode-Solution
	// 链接：https://leetcode-cn.com/problems/find-minimum-time-to-finish-all-jobs/solution/wan-cheng-suo-you-gong-zuo-de-zui-duan-s-hrhu/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}

// @lc code=end
