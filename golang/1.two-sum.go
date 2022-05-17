/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// 1. 需要存idx
// 2. 需要的是索引是target-idx
// @lc code=start
func twoSum(nums []int, target int) []int {
	mem := map[int]int{}
	for idx, num := range nums {
		if item, ok := mem[num]; ok {
			return []int{item, idx}
		}
		mem[target-num] = idx
	}
	return []int{}
}

// @lc code=end

