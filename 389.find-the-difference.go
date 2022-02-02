/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	ret := 0
	for i := 0; i < len(s); i++ {
		ret ^= int(s[i])
	}
	for i := 0; i < len(t); i++ {
		ret ^= int(t[i])
	}
	return byte(ret)

}

// @lc code=end

