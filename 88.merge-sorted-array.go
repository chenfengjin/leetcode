/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
// if you try to solve it from end to start
// it will be much easier than you have thought
package main

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m + n
	for l > 0 && n > 0 && m > 0 {
		if nums1[m-1] >= nums2[n-1] {
			nums1[l-1] = nums1[m-1]
			m -= 1
		} else {
			nums1[l-1] = nums2[n-1]
			n -= 1
		}
		l -= 1
	}
	if n > 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
}

// @lc code=end
