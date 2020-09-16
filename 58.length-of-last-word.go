/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.21%)
 * Total Accepted:    256.1K
 * Total Submissions: 795K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 */
import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	if len(s) == 0 {
		return 0
	}
	if strings.LastIndex(s, " ") < 0 {
		return len(s)
	}
	return len(s) - 1 - strings.LastIndex(s, " ")

}

