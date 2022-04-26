/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */
package main

import "fmt"

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	sign := (dividend > 0) == (divisor > 0)
	if dividend < 0 {
		dividend = dividend * -1
	}
	if divisor < 0 {
		divisor = divisor * -1
	}
	left := 0
	right := dividend
	for left < right {
		middle := (left + right) / 2
		if multiply(divisor, middle) < dividend {
			left = middle + 1
		} else {
			right = middle
		}
	}

	if multiply(left, divisor) > dividend {
		left = left - 1
	}
	if !sign {
		return left * -1
	}
	return left
}

func multiply(a, b int) int {
	n := 1
	cache := []int{}
	cur := a
	for n <= b {
		cache = append(cache, cur)
		cur = cur + cur
		n = n + n
	}
	// cache
	// 0 a*1
	// 1 a*2
	// 2 a*4
	// 3 a*8

	ret := 0
	i := 0
	for b > 0 {
		if b&1 == 1 {
			ret += cache[i]
		}
		i += 1
		b >>= 1
	}
	return ret
}

// @lc code=end

func main() {
	// fmt.Println(divide(6, 2))
	fmt.Printf("%b", 86)
}
