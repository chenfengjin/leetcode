/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

package main

// @lc code=start
func numTrees(n int) int {
	return numTreesWithCatalan(n)
}
func numTreesWithRecursion(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	ret := 0
	for i := 0; i < n; i++ {
		leftCount := numTrees(i)
		rightCount := numTrees((n - 1 - i))
		ret += leftCount * rightCount
	}
	return ret
}
func numTreesWithCatalan(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= (4*i - 2)
		ret /= (i + 1)
	}
	return ret
}

// @lc code=end

func main() {
	numTreesWithCatalan(10)
}
