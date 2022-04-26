/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}
	return false
}

// @lc code=end

