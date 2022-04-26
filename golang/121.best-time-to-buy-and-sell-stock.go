/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */
package main

import (
	"math"
)

// it is realy an "EASY" problem
//  you can never sell before buy
//  so just remember the slowest price untill now
//  and current price - slowest price means the profit you can get if you sell today

// @lc code=start
func maxProfit(prices []int) int {
	min := math.MaxInt32
	profit := 0
	for _, price := range prices {
		if min > price {
			min = price
		} else {
			if profit < price-min {
				profit = price - min
			}
		}
	}
	return profit
}

// @lc code=end

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}
