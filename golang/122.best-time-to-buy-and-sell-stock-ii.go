/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

//  dp[i] means max profit until the ith day
//  then
// 1. dp[0] = dp[0]
// dp[i] =  dp[i-1] 		if price[i] < price[i-1]
// 			price[i] - min	else
//
// TODO
// 官方题解很值得一看，无论是贪心还是动态规划
package main

import "fmt"

// @lc code=start
func maxProfit(prices []int) int {
	prices = append(prices, 0)
	min := prices[0]
	sum := 0
	for idx := 1; idx < len(prices); idx++ {
		price := prices[idx]
		if price < prices[idx-1] {
			sum += prices[idx-1] - min
			min = price
		}
	}
	return sum
}

// @lc code=end

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
