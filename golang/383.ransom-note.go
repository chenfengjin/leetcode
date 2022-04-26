/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */
package main

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {

	m := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		char := magazine[i]
		m[char] += 1
	}

	for i := 0; i < len(ransomNote); i++ {
		char := ransomNote[i]
		count, ok := m[char]
		if !ok {
			return false
		}
		if count == 0 {
			return false
		}
		m[char] -= 1
	}
	return true
}

// @lc code=end
