/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package main

import "fmt"

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLength := 200
	for _, str := range strs {
		if minLength > len(str) {
			minLength = len(str)
		}
	}
	findNotEqual := false
	//  important to initialize as -1
	longestPersifIndex := -1
	for i := 0; i < minLength && !findNotEqual; i++ {
		char := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != char {
				findNotEqual = true
				break
			}
		}
		if !findNotEqual {
			longestPersifIndex = i
		}
	}
	return strs[0][:longestPersifIndex+1]
}

// @lc code=end

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "dog", "dog"}))

}
