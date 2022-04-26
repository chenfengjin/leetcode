/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
//  dynamic programming
//  dp[i][j] means mininum number of opertions required to convert word[i] to word[j]
//  dp[i][j] =
//  if word1[i]=!word2[j] dp[i][j] is equal to  minimum number of operators of these three
//  1. dp[i-1][j-1] +1, replace a character
// 	2. dp[i][j-1] +1, add a character
//  3. dp[i-1][j] +1, delete a character

//  else if word1[i]== word2[j], dp[i] is determined by minimum number of these three
//  1. dp[i][j]  nothing to do
// 	2. dp[i][j-1] +1, add a character
//  3. dp[i-1][j] +1, delete a character
package main

import "fmt"

// dp[i][j] means edit distance of word1[0:i] word2[0:j]
// @lc code=start
func minDistance(word1 string, word2 string) int {
	// TODO  重新写一遍
	//  it will help to simplify code if add a common prefix to both word1 and word 2
	//  as it will make sure dp[0][0] == 0
	dp := [][]int{}
	for i := 0; i < len(word1)+1; i++ {
		dp = append(dp, make([]int, len(word2)+1))
	}

	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}

	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				min := dp[i-1][j-1]
				if min > dp[i-1][j]+1 {
					min = dp[i-1][j] + 1
				}
				if min > dp[i][j-1]+1 {
					min = dp[i-1][j-1] + 1
				}
				dp[i][j] = min
			} else {
				min := dp[i][j-1] + 1
				if min > dp[i-1][j]+1 {
					min = dp[i-1][j] + 1
				}
				if min > dp[i-1][j-1]+1 {
					min = dp[i-1][j-1] + 1
				}
				dp[i][j] = min
			}
		}
		// fmt.Println(dp[i-1])

	}
	return dp[len(word1)][len(word2)]
}

// @lc code=end
func main() {
	fmt.Println(minDistance("horse", "ros"))
}

//  good
//  goag
