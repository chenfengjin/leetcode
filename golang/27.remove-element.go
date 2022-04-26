/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {

	left := 0
	right := 0
	for ; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// @lc code=end

