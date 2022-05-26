/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] Maximal Square
 */
package main

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = minArray([]int{dp[i-1][j-1], dp[i][j-1], dp[i-1][j]}) + 1
				}
			}
			if maxSide < dp[i][j] {
				maxSide = dp[i][j]
			}
		}
	}
	return maxSide * maxSide
}
func minArray(nums []int) int {
	ret := nums[0]
	for _, num := range nums {
		if ret > num {
			ret = num
		}
	}
	return ret
}
func maxArray(nums []int) int {
	ret := nums[0]
	for _, num := range nums {
		if ret < num {
			ret = num
		}
	}
	return ret
}

// @lc code=end
