/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	x := 1 // 0
	y := 1 // 1
	z := 0
	for i := 2; i <= n; i++ {
		z = x + y
		x = y
		y = z
	}
	return z
}

// 初始条件
// 转移方程
// 终止条件
// package main
// func climbStairs(n int) int {

// 	memory := []int{1, 1}

// 	for i := 2; i <= n; i++ {
// 		current := memory[i-1] + memory[i-2]
// 		memory = append(memory, current)
// 	}
// 	return memory[n]
// }

// 递归内容
// 终止条件
// 边界条件

// var (
// 	memory = map[int]int{}
// )
// func climbStairs(n int) int {

// 	if n == 1 || n == 0 {
// 		return 1
// 	}
// 	if cached,ok:=memory[n];ok{
// 		return cached
// 	}

// 	result:= climbStairs(n-1) + climbStairs(n-2)
// 	memory [n]=result
// 	return result
// }

// @lc code=end

