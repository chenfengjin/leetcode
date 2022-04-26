/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
package main

// @lc code=start
func letterCombinations(digits string) []string {
	results := []string{}
	helper17(digits, "", 0, results)
	return results
}

var (
	mem = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
	}
)

//  pay attentin to escape analysi
func helper17(digits string, current string, pos int, results []string) {
	if pos == len(digits) {
		results = append(results, current)
	}

	for _, char := range mem[digits[pos]] {
		helper17(digits, current+char, pos+1, results)
	}
}

// @lc code=end

// 234
// 2 	a 	b 	c
// 3    d 	e 	f
// 4	h 	i 	j

// current ad
// pos 2

// adh adi adj
