/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */
//  binary search is useless as it is O(nlog(n)) time complexity and O(n) space complexity
// @lc code=start
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	sum := numbers[left] + numbers[right]
	for sum != target {
		if sum == target {
			break
		}
		if sum > target {
			right -= 1
		}
		if sum < target {
			left += 1
		}
		sum = numbers[left] + numbers[right]
	}
	return []int{left + 1, right + 1}
}

// @lc code=end

