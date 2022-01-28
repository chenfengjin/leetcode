/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */
package main

import "strings"

//  重点是对split 的函数定义，两个空格的情况处理
// @lc code=start
func reverseWords(s string) string {
	return reverseWordsStd(s)
}

func reverseWordsNonStd(s string) string {
	// TODO
	return ""

}
func reverseWordsStd(s string) string {
	words := strings.Split(s, " ")
	newWords := []string{}
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) == 0 {
			continue
		}
		newWords = append(newWords, words[i])
	}
	return strings.Join(newWords, " ")
}

// @lc code=end
