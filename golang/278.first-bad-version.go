/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
// 二分查找的几个关键点
// 1. left = middle+1
// 2. return right

func firstBadVersion(n int) int {
	left := 1
	right := n
	var middle int
	for left < right {
		middle = (left + right) / 2
		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return right
}

// @lc code=end

