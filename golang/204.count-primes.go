/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] Count Primes
 */
package main

import "fmt"

// @lc code=start
var (
	mask []bool
	dp   []int
)

func countPrimes(n int) int {
	return dp[n]

}

func init() {
	count := 5*1000000 + 1
	mask = make([]bool, count)
	dp = make([]int, count)
	dp[0] = 0
	mask[0], mask[1] = true, true
	i := 2
	for i < count {
		for i < count && mask[i] {
			i++
		}
		j := i + i
		for j < count {
			mask[j] = true
			j += i
		}
		i++
	}
	for i := 1; i < count; i++ {
		if !mask[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
}

// @lc code=end

func main() {
	fmt.Println(mask[0:10])
}
