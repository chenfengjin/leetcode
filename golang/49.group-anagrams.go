/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
package main

import "fmt"

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	table := map[string][]string{}
	for _, str := range strs {
		encode := encode(str)
		if _, ok := table[encode]; ok {
			table[encode] = append(table[str], str)
		} else {
			table[encode] = []string{str}
		}
	}
	ret := [][]string{}
	for _, list := range table {
		ret = append(ret, list)
	}
	return ret
}
func encode(str string) string {
	table := map[string]int{}
	for i := 0; i < len(str); i++ {
		if _, ok := table[string(str[i])]; ok {
			table[string(str[i])] += 1
		}
	}
	ret := ""
	for i := 0; i < 26; i++ {
		char := "abcedfeghijklmnokqrstuvwxyz"[i]
		if count, ok := table[string(char)]; ok {
			ret = fmt.Sprintf("%s%d", char, count)
		}
	}
	return ret
}

// @lc code=end
