/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
// a method to optimize performance is store index of character in hash table
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	longestStr := ""
	// a01234bcbefg
	left := 0
	right := 1
	mem := map[str]int{}
	for ; right < len(s); right++ {
		if idx, ok := mem[s[right]]; ok {
			if maxLength < right-idx{
				maxLength = right-idx
				longestStr = s[idx,right]
			}
		}
	}
}

// @lc code=end

