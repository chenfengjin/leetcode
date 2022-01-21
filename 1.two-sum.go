/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	diff := map[int]int{}
	for idx, num := range nums {
		if value, ok := diff[target-num]; ok {
			return []int{value, idx}
		}
		diff[num] = idx
	}
	return []int{}
}

// @lc code=end

