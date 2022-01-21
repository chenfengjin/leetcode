/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	old := x
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x = x / 10
	}
	return old == reversed
}

// @lc code=end

// func main() {
// 	isPalindrome(121)
// }
