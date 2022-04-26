/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
package main

import "fmt"

// @lc code=start
// 快速幂算法
//  无须处理 x 的符号
// 需要特殊处理 n 的符号
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	sign := 1
	if n < 0 {
		sign = -1
		n = -1 * n
	}

	cache := []float64{}
	m := 1
	current := x
	for m <= n {
		cache = append(cache, current)
		current = current * current
		m = m * 2
	}
	result := 1.0
	bit := 0
	for n > 0 {
		if n&0x1 == 1 {
			result *= cache[bit]
		}
		bit += 1
		n = n >> 1
	}

	if sign == -1 {
		return 1 / result
	}
	return result
}

// @lc code=end

func main() {
	fmt.Println(myPow(2.0, -2))
}
