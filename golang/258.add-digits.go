/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start
func addDigits(num int) int {

	for num > 9 {
		new := 0
		for num > 0 {
			new += num % 10
			num /= 10
		}
		num = new
	}
	return num
}

// @lc code=end

