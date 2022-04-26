/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	// 1   1
	// 11   2
	// 21   3
	// 1211     4
	// 111221		5
	// 312211 6
	// s=s+"#"
// 哨兵机制
	// 1	1	1	2	2	1	#
	// ^			^		
	// |			|
	// left 		right
	
	for right<len(s) {
		left := 0
		right := 1
		if s[right] == s[right-1] {
			continue
		} else {

		}
		right += 1
	}
	// 312211
}

// @lc code=end

